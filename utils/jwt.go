package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenereateTokens(userId int) string {
	secretkey := os.Getenv("APP_SECRET")

	accessClaims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(1 * time.Hour).Unix(),
		"iat":    time.Now().Unix(),
	}
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	token, _ := generateToken.SignedString([]byte(secretkey))

	return token
}
