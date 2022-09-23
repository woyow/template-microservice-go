package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/woyow/example-module/config"
	"github.com/woyow/example-module/internal/queue/nats"
	"github.com/woyow/example-module/internal/queue/redismq"
	"github.com/woyow/example-module/internal/service"
	"github.com/woyow/example-module/internal/storage"
	"github.com/woyow/example-module/internal/storage/psql"
	"github.com/woyow/example-module/internal/storage/redis"
	"github.com/woyow/example-module/internal/transport/grpc"
	"github.com/woyow/example-module/internal/transport/http"
	httpHandler "github.com/woyow/example-module/internal/transport/http/handler"
)

func Run(cfg *config.Config) {

	// Logger initialization
	logger := NewLogger(&cfg.Log)
	logger.Info("Logger has been initialized")

	// Postgresql migrations
	migrateConfig := NewMigrateConfig(&cfg.App, &cfg.PG, logger)
	migrateConfig.Migrate()

	// Postgresql database initialization
	psqlDB, err := psql.NewPsqlDB(&cfg.PG)
	if err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}
	logger.Info("Postgresql db has been initialized")
	defer psqlDB.Close() // Close database connection after return of function

	// Redis cache initialization
	redisDB := redis.NewRedisClient(&cfg.Redis)
	logger.Info("Redis db has been initialized")
	defer redisDB.Close() // Close redis connection after return of function

	// Dependency injections
	storages := storage.NewStorage(psqlDB, redisDB, logger)
	services := service.NewService(storages, cfg, logger)
	redisMQ := redismq.NewRedisMQ(&cfg.RedisMQ, storages, logger)
	go func() {
		if err := redisMQ.Run(context.Background()); err != nil {
			logger.Fatalf("error occurred while running redismq: %s", err.Error())
		}
	}()

	// Nats
	natsMQ, err := nats.NewNats(&cfg.Nats, storages, logger)
	if err != nil {
		logger.Error(err)
	}
	go func() {
		natsMQ.Run()
	}()

	handlers := httpHandler.NewHandler(&cfg.App.HttpHandler, services, redisMQ, natsMQ, logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Run HTTP server
	httpServer := http.NewServer(&cfg.HTTP, handlers.Init(), logger)
	logger.Info("Initialize http server")
	go func() {
		if err := httpServer.Run(ctx); err != nil {
			logger.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()

	// Run GRPC server
	grpcServer := grpc.NewServer(&cfg.GRPC, logger)
	logger.Info("Initialize grpc server")
	go func() {
		if err := grpcServer.Run(); err != nil {
			logger.Fatalf("error occurred while running grpc server: %s", err.Error())
		}
	}()

	logger.Printf("%s:%s started\n", cfg.App.Name, cfg.App.Version)

	// Handle signals for shutdown services
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Printf("%s:%s shutdown\n", cfg.App.Name, cfg.App.Version)

	// Shutdown services
	natsMQ.Shutdown()
	redisMQ.Shutdown(context.Background())

	if err := httpServer.Shutdown(context.Background()); err != nil {
		logger.Errorf("error occurred on http server shutting down: %s", err.Error())
	}

	grpcServer.Shutdown()
}