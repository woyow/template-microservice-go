package app

import (
	"github.com/woyow/example-module/config"
	
	"github.com/sirupsen/logrus"
)

func NewLogger(cfg *config.Log) *logrus.Logger {
	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: cfg.DisableTimestamp,
		FullTimestamp: cfg.FullTimestamp,
	})

	// Possible logLevel value: "panic", "fatal", "error", "warn" or "warning", "info", "debug", "trace"
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		logger.WithError(err).WithField("level", cfg.Level).Warn("cannot parse a logging level")
	} else {
		logger.SetLevel(level)
	}

	return logger
}