package db

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// New initializes and returns a postgres database connection based on the provided configuration.
func New(cfg Config) (*sql.DB, *gorm.DB, error) {
	// connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode)

	// open connection
	pg, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, nil, err
	}

	pg.SetMaxOpenConns(25)
	pg.SetMaxIdleConns(25)
	pg.SetConnMaxLifetime(20 * time.Minute)
	pg.SetConnMaxIdleTime(5 * time.Minute)

	// ping to verify connection
	if err := pg.Ping(); err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return pg, gormDB, nil
}
