package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nicewook/authjwt/handler"
)

type jwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

const (
	uploadedDir = "uploaded"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// port setting. default is 80 and 443
	var ports string
	flag.StringVar(&ports, "port", "80:443", "set http/https port. {http-port}:{https-port}")
	flag.Parse()

	p := strings.Split(ports, ":")
	httpPort := ":" + p[0]
	httpsPort := ":" + p[1]

	log.Printf("http port is %v, https port is %v\n", httpPort, httpsPort)

	// make uploaded folder. MkdirAll will not do anything if directory already exist
	if err := os.MkdirAll(uploadedDir, 0755); err != nil {
		log.Fatal(err)
	}

	// echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", handler.Login)
	e.GET("/", handler.Accessible)

	// Restricted group
	r := e.Group("/v1")
	{
		// Configure middleware with the custom claims type
		secretKey := "should be env secret" // 나중에는 환경변수에 넣어주자. 클라이언트도 환경변수에서 읽게 하자
		config := middleware.JWTConfig{
			Claims:     &jwtCustomClaims{},
			SigningKey: []byte(secretKey),
		}
		r.Use(middleware.JWTWithConfig(config))
		r.POST("/signup", handler.SignUp)
		r.POST("/upload", handler.Upload)
	}

	// 목데이터로 테스트
	e.GET("/api/getlist", handler.MockData(), middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "cookie:access-token",
	}))

	go func() {
		e.Logger.Fatal(e.Start(httpPort))
	}()
	e.Logger.Fatal(e.StartTLS(httpsPort, "server.crt", "server.key"))

}
