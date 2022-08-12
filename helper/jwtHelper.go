package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// https://echo.labstack.com/cookbook/jwt/

type jwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func CreateJWT(name string) (string, error) {

	secretkey := "should be env secret" // 나중에는 환경변수에 넣어주자. 클라이언트도 환경변수에서 읽게 하자
	// Set custom claims
	claims := &jwtCustomClaims{
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(secretkey))
}
