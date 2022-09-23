package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type Cache struct {

}

// NewCache returns redis cache
func NewCache(rdb *redis.Client, logger *logrus.Logger) *Cache {
	return &Cache{

	}
}