package storage

import (
	"context"
	"go-bookman-app/utils"
	"log"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func ConnectRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     utils.GetConfig("REDIS_ADDRESS"),
		Password: utils.GetConfig("REDIS_PASSWORD"),
		DB:       0,
	})

	if err := RDB.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("error when connecting to redis: %s\n", err)
	}

	log.Println("redis connected")
}
