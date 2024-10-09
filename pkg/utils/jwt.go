package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	SecretKey = "asfa313sdg##F"
	TokenTTL  = 15 * time.Hour
)

func ParseJWT(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(SecretKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	if token.Valid {
		return claims["id"], nil
	}

	return nil, nil
}

func GenerateJWT(userID string, userName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   userID,
		"name": userName,
		"exp":  time.Now().Add(TokenTTL).Unix(),
	})

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
