package psql

import (
	"github.com/woyow/example-module/config"

	_"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"context"
)

const (
	exampleTable = "example_table"
)

func NewPsqlDB(cfg *config.PG) (*pgxpool.Pool, error) {

	// Configure database host and port
	if cfg.PgBouncer.Enable == true {
		cfg.Host = cfg.PgBouncer.Host
		cfg.Port = cfg.PgBouncer.Port
	}

	pgProto := "postgres://"
	dbSource := pgProto + cfg.Username + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.DBName

	pool, err := pgxpool.Connect(context.Background(), dbSource)
	if err != nil {
		return nil, err
	}

	pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return pool, nil
}