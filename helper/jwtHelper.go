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

func CreateJWT(name string, secretKey string) (string, error) {

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
	return token.SignedString([]byte(secretKey))
}
