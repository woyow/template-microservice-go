# /internal/storage/psql

Directory for working with postgresql driver

## Example
Open `psql.go` file. Add new interface `Something`:
```go
package psql

import (
	"{{SERVICE_NAME}}/internal/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Something interface {
	CreateSomething(something *entity.CreateSomethingReq) (string, error)
	GetSomethingByID(somethingID string) (*entity.GetSomethingResp, error)
}

type PostgresStorage struct {
	Something *SomethingStorage
}

func NewPostgresStorage(db *pgxpool.Pool) *PostgresStorage {
	return &PostgresStorage{
        Something: NewSomethingStorage(db),
	}
}
```

Create file `something.go`:
```go
import(
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"

	"{{SERVICE_NAME}}/internal/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
    somethingTable = "something"
)

type SomethingStorage struct {
	db *pgxpool.Pool
}

func NewSomethingStorage(db *pgxpool.Pool) *SomethingStorage {
	return &SomethingStorage{db: db}
}

func (s *SomethingStorage) CreateSomething(something *entity.CreateSomethingReq) (string, error) {
    createSomethingQuery := fmt.Sprintf("INSERT INTO %s " +
    "(id, created_at, name, city, age) " +
    "VALUES ($1, $2, $3, $4, $5) " +
    "RETURNING id", somethingTable)
    
    var newSomethingID string
    
    if err := s.db.QueryRow(context.Background(), createSomethingQuery,
    something.ID,
    something.CreatedAt,
    something.Name,
    something.City,
    something.Age).Scan(&newSomethingID); err != nil {
    log.Debug("psql: create something query error: ", err.Error())
    return nil, err
    }

    log.Debug("psql: new something id: ", newSomethingID)
    
	return &newSomethingID, nil
}

func (s *SomethingStorage) GetSomethingByID(somethingID string) (*entity.GetSomethingResp, error) {
	var getSomething entity.GetSomethingResp
    
    getSomethingByIDQuery := fmt.Sprintf("SELECT " +
    "id, created_at, name, city, age " +
    "FROM %s WHERE id=$1;", somethingTable)
    
    if err := s.db.QueryRow(context.Background(), getSomethingByIDQuery, somethingID).Scan(
    &getSomething.ID,
    &getSomething.CreatedAt,
    &getSomething.Name,
    &getSomething.City,
    &getSomething.Age); err != nil {
    log.Debug("psql: get something query error: ", err.Error())
    return nil, err
    }
    
    log.Debug("psql: get something by id: ", getSomething)
    
    return &getSomething, nil
}

```