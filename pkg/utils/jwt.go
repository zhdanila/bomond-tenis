package utils

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	SecretKey = "asfa313sdg##F"
	TokenTTL  = 15 * time.Minute
)

func CheckJWT(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(SecretKey), nil
	})
	if err != nil {
		return err
	}

	if token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func GetClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
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

func JWTBlacklistMiddleware(tokenString string, redisClient *redis.Client) error {
	val, err := redisClient.Get(context.Background(), tokenString).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}

	if val == "blacklisted" {
		return errors.New("token is blacklisted")
	}

	return nil
}
