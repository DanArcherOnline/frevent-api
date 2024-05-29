package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "secret"

func GenerateToken(userId int64, email string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
