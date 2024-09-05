package utils

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func NewRedisDB(ctx context.Context, host, port, password string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
