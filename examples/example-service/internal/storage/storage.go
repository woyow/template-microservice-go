package storage

import (
	"github.com/woyow/example-module/internal/storage/psql"
	"github.com/woyow/example-module/internal/storage/redis"

    "github.com/sirupsen/logrus"
	"github.com/jackc/pgx/v4/pgxpool"
	rds "github.com/go-redis/redis/v8"
)

type Storage struct {
	Psql psql.Storage
	Redis redis.Cache
}

func NewStorage(psqlDB *pgxpool.Pool, redisDB *rds.Client, logger *logrus.Logger) *Storage {
	return &Storage{
		Psql: *psql.NewStorage(psqlDB, logger),
		Redis: *redis.NewCache(redisDB, logger),
	}
}