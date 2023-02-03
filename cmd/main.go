package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/srselivan/user-balance-microservice/internal/app/server"
)

func GetViperConfig(configType string, path string) error {
	viper.SetConfigType(configType)
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := GetViperConfig("toml", "./configs"); err != nil {
		logrus.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatal()
	}
	viper.Set("database.password", os.Getenv("DB_PASSWORD"))

	s := server.New(viper.GetViper())
	if err := s.Start(); err != nil {
		logrus.Fatal(err)
	}
}
