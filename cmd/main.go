package main

import (
	"log"

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
	log.Println(viper.Get("server.port"))

	return nil
}

func main() {
	if err := GetViperConfig("toml", "./configs"); err != nil {
		logrus.Fatal(err)
	}

	s := server.New(viper.GetViper())
	if err := s.Start(); err != nil {
		logrus.Fatal(err)
	}
}
