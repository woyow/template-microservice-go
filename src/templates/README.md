# Run application

## 0. Go to root project folder
```bash
cd ./{{SERVICE_NAME}}
```

## 1. Create .env file
```bash 
cp ./.env.example ./.env
```

## 2. Fill .env file
```bash
nano ./.env
```
### .env example:
```nano
APP_ENV: "local"

REDIS_USERNAME: ""
REDIS_PASSWORD: ""

PG_USERNAME: "postgres"
PG_PASSWORD: "postgres"

JWT_SECRET: "extra-secret"
```

## 3. Download deps
```bash
go mod tidy
```

## 4. Create database for microservice
```bash
psql -U postgres
CREATE DATABASE IF NOT EXISTS local_{{SERVICE_NAME}}_v1;
```

## 5. Run application
```bash
go run ./cmd/{{SERVICE_NAME}}/main.go
```