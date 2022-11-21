package database

import (
	"database/sql"
	"sync"

	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type database struct {
	store  *sql.DB
	config *Config
	logger *logrus.Logger
}

var instance *database
var once sync.Once

func New(config *Config) *database {
	once.Do(func() {
		instance = &database{
			store:  &sql.DB{},
			config: config,
			logger: logrus.New(),
		}
	})

	return instance
}

func (db *database) ConnectToDB() error {
	cfg := mysql.Config{
		User:   db.config.User,
		Passwd: db.config.Passwd,
		Net:    db.config.Net,
		Addr:   db.config.Addr,
		DBName: db.config.DBName,
	}

	var err error
	db.store, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}

	err = db.store.Ping()
	if err != nil {
		return err
	}

	db.logger.Infof("Connect to database %s, on addr %s, with timeout %v", cfg.DBName, cfg.Addr, cfg.Timeout)
	return nil
}

func (db *database) CloseConnection() error {
	db.logger.Info("Close connection to database")
	return db.store.Close()
}

func (db *database) configureLogger() error {
	level, err := logrus.ParseLevel(db.config.LogLevel)
	if err != nil {
		return err
	}

	logrus.SetLevel(level)
	return nil
}
