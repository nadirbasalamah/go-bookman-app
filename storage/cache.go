package storage

import (
	"context"
	"go-bookman-app/utils"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     utils.GetConfig("REDIS_ADDRESS"),
		Password: utils.GetConfig("REDIS_PASSWORD"),
		DB:       0,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
