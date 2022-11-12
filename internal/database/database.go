package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type database struct {
	db     *sql.DB
	config *mysql.Config
	logger *logrus.Logger
}

func ConfigureDatabase(config *mysql.Config) error {
	return nil
}
