package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func InitDB(cfg Config) error {
	if cfg.Host == "" {
		cfg.Host = "localhost"
	}

	if cfg.Port == "" {
		cfg.Port = "5432"
	}

	if cfg.SSLMode == "" {
		cfg.SSLMode = "disable"
	}

	if cfg.User == "" {
		return errors.New("DB_USER is required")
	}

	if cfg.Password == "" {
		return errors.New("DB_PASSWORD is required")
	}

	if cfg.DBName == "" {
		return errors.New("DB_NAME is required")
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("ошибка ping БД: %w", err)
	}

	log.Println("подключение к БД установлено")
	return nil
}
