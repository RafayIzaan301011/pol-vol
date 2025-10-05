package db

import (
	"database/sql"
	"fmt"
	"time"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func New(cfg Config) (*sql.DB, error) {
	// connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode)

	// open connection
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxLifetime(20 * time.Minute)
	conn.SetConnMaxIdleTime(5 * time.Minute)

	// ping to verify connection
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
