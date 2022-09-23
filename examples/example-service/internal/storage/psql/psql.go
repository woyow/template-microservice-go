package psql

import (
    "github.com/jackc/pgx/v4/pgxpool"
    "github.com/sirupsen/logrus"
)

type Storage struct {

}

// NewStorage returns postgresql storage
func NewStorage(db *pgxpool.Pool, logger *logrus.Logger) *Storage {
	return &Storage{

	}
}