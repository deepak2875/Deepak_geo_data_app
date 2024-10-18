package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your_secret_key")

func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(jwtKey)
}
