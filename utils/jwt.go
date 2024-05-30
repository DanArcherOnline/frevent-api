package utils

import (
	"errors"
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

func VerifyToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if !isHMACALgorithm(token) {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil || !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	return parsedToken, nil
}

func isHMACALgorithm(token *jwt.Token) bool {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	return ok
}

func GetUserIDFromToken(token *jwt.Token) (int64, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("unexpected claims type")
	}
	userID, ok := claims["userId"].(float64)
	if !ok {
		return 0, errors.New("type assertion failed")
	}
	userID64 := int64(userID)
	return userID64, nil
}
