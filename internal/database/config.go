package database

import "time"

type Config struct {
	User     string        `toml:"user"`
	Passwd   string        `toml:"passwd"`
	Net      string        `toml:"net"`
	Addr     string        `toml:"addr"`
	DBName   string        `toml:"dBName"`
	LogLevel string        `toml:"database_log_level"`
	Timeout  time.Duration `toml:"timeout"`
}

func NewConfig() *Config {
	return &Config{
		User:     "root",
		Passwd:   "root",
		Net:      "tcp",
		Addr:     "127.0.0.1:3306",
		DBName:   "mysql",
		LogLevel: "debug",
		Timeout:  120 * time.Second,
	}
}
