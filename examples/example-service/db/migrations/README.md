# Migrations for database

## Tool
https://github.com/golang-migrate/migrate

## Create up&down migration
```bash
migrate create -ext sql -dir ./db/migrations/ migration_name
```