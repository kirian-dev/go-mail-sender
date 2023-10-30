package middleware

import (
	"go-mail-sender/config"

	"github.com/sirupsen/logrus"
)

// Middleware manager
type MiddlewareManager struct {
	cfg *config.Config
	log *logrus.Logger
}

func NewMiddlewareManager(cfg *config.Config, log *logrus.Logger) *MiddlewareManager {
	return &MiddlewareManager{cfg: cfg, log: log}
}
