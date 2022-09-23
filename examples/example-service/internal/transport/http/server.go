package http

import (
	"github.com/woyow/example-module/config"

	"github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"

	"context"
	"errors"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	log *logrus.Logger
}

// NewServer returns http server
func NewServer(cfg *config.HTTP, handler http.Handler, logger *logrus.Logger) *Server {

	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.Port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadHeaderTimeout: 1 * time.Second,
			IdleTimeout: 	10 * time.Second,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
		log: logger,
	}
}

func (s *Server) Run(ctx context.Context) error {
	logger := log.Ctx(ctx)
	logger.Info().Str("Http server", "Start listening")
	defer func() {
		logger.Info().Str("Http server", "Stop listening")
	}()

	if err := s.httpServer.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		logger.Error().Err(err).Str("Http server", "Failed start listening")
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}