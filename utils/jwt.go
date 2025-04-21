package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(data any, exp time.Duration, secret string) (string, error) {
	claims := jwt.MapClaims{
		"data": data,
		"exp":  time.Now().Add(time.Hour * exp).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func ValidateToken(token string, secret string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKey
		}

		return secret, nil
	})
}
