package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	"log"
	"os"
	"time"
)

const (
	configDir = "config"
)

type (
	// Config - Main config for application
	Config struct {
		App
		Middleware
		Log
		PG
		Redis
		Nats
		RedisMQ
	}

	// HTTP - Http server configuration
	HTTP struct {
		HttpHandler
		Port string
	}

	// HttpHandler - Http handler configuration
	HttpHandler struct {
		CORS
	}

	// CORS - Http cors configuration
	CORS struct {
		AllowCredentials bool
		AllowAllOrigins bool
		AllowMethods []string
		AllowHeaders []string
		AllowOrigins []string
	}

	// GRPC - Grpc server configuration
	GRPC struct {
		Proto string
		Port string
	}

	// App - Application configuration
	App struct {
		Env 	string
		Name    string
		Version string
		HTTP
		GRPC
	}

	// JWT - Json web token configuration
	JWT struct {
		Secret 		string
	}

	// Middleware - Middleware configuration
	Middleware struct {
		JWT
	}

	// Log - Logger configuration
	Log struct {
		Level string
		DisableTimestamp bool
		FullTimestamp bool
	}

	// PgBouncer - Pgbouncer configuration
	PgBouncer struct {
		Enable bool
		Host string
		Port string
	}

	// PG - Postgresql database
	PG struct {
		Host string
		Port string
		Username string
		DBName string
		SSLMode string
		Password string
		PgBouncer
	}

	// Redis - Cache
	Redis struct {
		Host     string
		Port     string
		DB       int
		Password string
	}

	// Nats - Message queue
	Nats struct {
		Host 	string
		Port 	string
	}

	// RedisMQCleaner - Cleaner configuration
	RedisMQCleaner struct {
		CleanPeriod 	time.Duration
	}

	// RedisMQConsumer - Consumer configuration
	RedisMQConsumer struct {
		NumberForQueue 		int
	}

	// RedisMQ - Pub/Sub message broker
	RedisMQ struct {
		Host	 			string
		Port 				string
		Username 			string
		Password 			string
		DB		 			int
		PoolSize 			int
		RedisMQCleaner
		RedisMQConsumer
	}
)

type EnvParameters struct {
	AppEnv 			string
	JWTSecret 		string
	PGUsername 		string
	PGPassword 		string
	RedisUsername 	string
	RedisPassword 	string
}

func ReadEnvParameters() (*EnvParameters, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s\n", err.Error())
		return nil, err
	}

	return &EnvParameters{
		AppEnv: os.Getenv("APP_ENV"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		PGUsername: os.Getenv("PG_USERNAME"),
		PGPassword: os.Getenv("PG_PASSWORD"),
		RedisUsername: os.Getenv("REDIS_USERNAME"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
	}, nil
}

func ReadConfig(configFileName string) error {
	viper.AddConfigPath(configDir)
	viper.SetConfigName(configFileName)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("config error: %s", err.Error())
		return err
	}

	return nil
}

func SetConfigVariables(envParams *EnvParameters) *Config {
	return &Config{
		App: App{
			Env: envParams.AppEnv,
			Name: viper.GetString("app.name"),
			Version: viper.GetString("app.version"),
			HTTP: HTTP{
				HttpHandler: HttpHandler{
					CORS: CORS{
						AllowCredentials: viper.GetBool("app.http.handler.cors.allow_credentials"),
						AllowAllOrigins: viper.GetBool("app.http.handler.cors.allow_all_origins"),
						AllowMethods: viper.GetStringSlice("app.http.handler.cors.allow_methods"),
						AllowHeaders: viper.GetStringSlice("app.http.handler.cors.allow_headers"),
						AllowOrigins: viper.GetStringSlice("app.http.handler.cors.allow_origins"),
					},
				},
				Port: viper.GetString("app.http.port"),
			},
			GRPC: GRPC{
				Proto: viper.GetString("app.grpc.proto"),
				Port: viper.GetString("app.grpc.port"),
			},
		},
		Middleware: Middleware{
			JWT: JWT{
				Secret: envParams.JWTSecret,
			},
		},
		Log: Log{
			Level: viper.GetString("logger.log_level"),
			DisableTimestamp: viper.GetBool("logger.disable_timestamp"),
			FullTimestamp: viper.GetBool("logger.full_timestamp"),
		},
		PG: PG{
			Username: envParams.PGUsername,
			Password: envParams.PGPassword,
			Host: viper.GetString("postgres.host"),
			Port: viper.GetString("postgres.port"),
			DBName: viper.GetString("postgres.dbname"),
			SSLMode: viper.GetString("postgres.sslmode"),
			PgBouncer: PgBouncer{
				Enable: viper.GetBool("postgres.pgbouncer.enable"),
				Host: viper.GetString("postgres.pgbouncer.host"),
				Port: viper.GetString("postgres.pgbouncer.port"),
			},
		},
		Redis: Redis{
			Host: viper.GetString("redis.host"),
			Port: viper.GetString("redis.port"),
			DB: viper.GetInt("redis.db"),
			Password: envParams.RedisPassword,
		},
		Nats: Nats{
			Host: viper.GetString("nats.host"),
			Port: viper.GetString("nats.port"),
		},
		RedisMQ: RedisMQ{
			Host: viper.GetString("redismq.host"),
			Port: viper.GetString("redismq.port"),
			Username: envParams.RedisUsername,
			Password: envParams.RedisPassword,
			DB: viper.GetInt("redismq.db"),
			PoolSize: viper.GetInt("redismq.pool_size"),
			RedisMQCleaner: RedisMQCleaner{
				CleanPeriod: viper.GetDuration("redismq.cleaner.clean_period"),
			},
			RedisMQConsumer: RedisMQConsumer{
				NumberForQueue: viper.GetInt("redismq.consumer.number_for_queue"),
			},
		},
	}
}

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	// Initialization of .env file and read env parameters
	envParams, err := ReadEnvParameters()
	if err != nil {
		return nil, err
	}
	log.Printf("config env: %s\n\n", envParams.AppEnv)

	// Set config file and read config
	if err := ReadConfig(envParams.AppEnv); err != nil {
		return nil, err
	}

	return SetConfigVariables(envParams), nil
}