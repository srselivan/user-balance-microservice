package database

import (
	"database/sql"
	"sync"

	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/srselivan/user-balance-microservice/internal/server"
)

type database struct {
	db     *sql.DB
	config *server.Config
	logger *logrus.Logger
}

var instance *database
var once sync.Once

func New(config *server.Config) *database {
	once.Do(func() {
		instance = &database{
			db:     &sql.DB{},
			config: config,
			logger: logrus.New(),
		}
	})

	return instance
}

func (db *database) ConnectToDB() error {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "mysql",
	}

	var err error
	db.db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer db.db.Close()

	err = db.db.Ping()
	if err != nil {
		return err
	}

	db.logger.Info("Connect to database")
	return nil
}

func (db *database) CloseConnection() error {
	return db.db.Close()
}

func (db *database) configureLogger() error {
	level, err := logrus.ParseLevel(db.config.LogLevel)
	if err != nil {
		return err
	}

	logrus.SetLevel(level)
	return nil
}
