package db

import (
	"go.uber.org/zap"
)

// Config stores the configuration for the database.
type Config struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	Logger   *zap.Logger
}
