package main

import (
	"log"
	"os"

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

	// godotenv는 로컬 개발환경에서 .env를 통해 환경변수를 읽어올 때 쓰는 모듈이다.
	// 프로덕션 환경에서는 필요하지 않음.
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	log.SetFlags(log.LstdFlags | log.Lshortfile)

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

	e.Logger.Fatal(e.Start(":1323"))
}
