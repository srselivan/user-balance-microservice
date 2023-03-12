package main

import (
	"context"
	"github.com/srselivan/user-balance-microservice/internal/app/handler"
	"github.com/srselivan/user-balance-microservice/internal/app/repository"
	"github.com/srselivan/user-balance-microservice/internal/app/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/srselivan/user-balance-microservice/internal/app/server"
)

const shutdownTimeout = 10 * time.Second

func GetViperConfig(configType string, path string) error {
	viper.SetConfigType(configType)
	viper.AddConfigPath(path)
	return viper.ReadInConfig()
}

func main() {
	err := GetViperConfig("toml", "./configs")
	if err != nil {
		logrus.Fatal(err)
	}

	err = godotenv.Load()
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := repository.NewDB(repository.Config{
		Port:     viper.GetInt("database.port"),
		Host:     viper.GetString("database.host"),
		Username: viper.GetString("database.user"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("database.dbname"),
		SSLMode:  viper.GetString("database.sslmode"),
	})
	if err != nil {
		logrus.Fatal(err)
	}

	repo := repository.New(db)
	services := service.New(repo)
	handlers := handler.New(services)

	s := new(server.Server)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		err = s.Run(viper.GetString("server.port"), handlers)
		if err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("server run: %v", err)
		}
	}()

	<-ctx.Done()
	logrus.Info("server shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err = s.Shutdown(shutdownCtx)
	if err != nil {
		logrus.Fatalf("shutdown: %v", err)
	}

	err = db.Close()
	if err != nil {
		logrus.Fatalf("database: %v", err)
	}

	logrus.Info("server shutdown")
}
