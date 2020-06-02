package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// Conn represents a connection to the database.
type Conn struct {
	*gorm.DB
	logger *zap.Logger
}

// New creates a new database connection and initializes it.
func New(cfg *Config) (*Conn, error) {
	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
			cfg.Host,
			cfg.Port,
			cfg.Name,
			cfg.User,
			cfg.Password,
		),
	)
	if err != nil {
		return nil, err
	}
	return &Conn{
		DB:     db,
		logger: cfg.Logger.Named("db"),
	}, nil
}
