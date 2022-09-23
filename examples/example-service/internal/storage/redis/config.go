package redis

import (
	"github.com/woyow/example-module/config"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(cfg *config.Redis) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return redisClient
}