package server

type Config struct {
	Port     string `toml:"port"`
	LogLevel string `toml:"server_log_level"`
}

func NewConfig() *Config {
	return &Config{
		Port:     ":8080",
		LogLevel: "debug",
	}
}
