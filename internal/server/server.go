package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/srselivan/user-balance-microservice/internal/handlers"
)

type server struct {
	config *Config
	router *mux.Router
	logger *logrus.Logger
}

func New(config *Config) *server {
	return &server{
		config: config,
		router: mux.NewRouter(),
		logger: logrus.New(),
	}
}

func (s *server) Run() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Infof("Starting server on %v port", s.config.Port)
	s.router.HandleFunc("/", handlers.HandleHello)

	return http.ListenAndServe(s.config.Port, s.router)
}

func (s *server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	logrus.SetLevel(level)

	return nil
}
