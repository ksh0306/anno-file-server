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

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var ports string
	flag.StringVar(&ports, "port", "8080:8443", "set http/https port. {http-port}:{https-port}")

	p := strings.Split(ports, ":")
	httpPort := ":" + p[0]
	httpsPort := ":" + p[1]

	log.Printf("http port is %v, https port is %v\n", httpPort, httpsPort)

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

	e.Logger.Fatal(e.Start(httpPort))
}
