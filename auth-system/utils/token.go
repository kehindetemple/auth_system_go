package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID string, key string, expTime int) (string, error) {

	claims := jwt.MapClaims{}

	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Duration(expTime) * time.Hour).Unix()

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}