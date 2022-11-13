package server

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/srselivan/user-balance-microservice/internal/handlers"
)

type server struct {
	config *Config
	router *mux.Router
	logger *logrus.Logger
}

var instance *server
var once sync.Once

func New(config *Config) *server {
	once.Do(func() {
		instance = &server{
			config: config,
			router: mux.NewRouter(),
			logger: logrus.New(),
		}
	})

	return instance
}

func (s *server) Start() error {
	err := s.configureLogger()
	if err != nil {
		s.logger.Info("Can't configure logger. Logger now configure by default settings")
	}

	s.configureRouter()

	s.logger.Infof("Starting server on %s port", s.config.Port)
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

func (s *server) configureRouter() {
	s.router.HandleFunc("/health", handlers.Health().ServeHTTP).Methods("GET")
}
