package repository

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", GetConnectionString(cfg))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetConnectionString(cfg Config) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode,
	)
}
