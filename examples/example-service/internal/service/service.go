package service

import (
    "github.com/woyow/example-module/config"
	"github.com/woyow/example-module/internal/storage"

	"github.com/sirupsen/logrus"
)


type Service struct {

}

func NewService(storage *storage.Storage, cfg *config.Config, logger *logrus.Logger) *Service {
	return &Service{

	}
}
