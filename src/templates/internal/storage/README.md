# /internal/storage

Directory for logic to work with databases

For example, we have different methods for postgres, mysql databases

File `storage.go`:
```go
import (
	"your-module-name/internal/storage/psql"
	"your-module-name/internal/storage/mysql"
)

type Storage struct {
	Psql psql.PostgresStorage
	Mysql mysql.MysqlStorage
}

func NewStorage(psqlDB *<pgDriver>, mysqlDB *<mysqlDriver>) *Storage {
    return &Storage{
        Psql: *psql.NewPostgresStorage(psqlDB),
        Mysql: *mysql.NewMysqlStorage(mysqlDB),
    }
}
```