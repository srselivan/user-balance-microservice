package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/srselivan/user-balance-microservice/internal/app/handlers"
	"github.com/srselivan/user-balance-microservice/internal/app/store"
)

type Server struct {
	config *viper.Viper
	router *mux.Router
	store  *store.Store
}

func New(config *viper.Viper) *Server {
	return &Server{
		config: config,
		router: mux.NewRouter(),
		store:  store.New(),
	}
}

func (s *Server) Start() error {
	s.configureRouter()

	connStr := s.constructConnString()
	s.store.ConnectToDB(connStr)

	logrus.Infof("Starting server on %d port", s.config.Get("server.port"))

	return http.ListenAndServe(fmt.Sprintf(":%d", s.config.Get("server.port")), s.router)
}

func (s *Server) configureRouter() {
	handler := handlers.New(s.store)

	s.router.HandleFunc("/health", handlers.Health().ServeHTTP).Methods("GET")

	s.router.HandleFunc("/balance", handler.GetBalance().ServeHTTP).Methods("GET")
	s.router.HandleFunc("/balance", handler.AppendBalance().ServeHTTP).Methods("PUT")

	s.router.HandleFunc("/reserve", handler.ReserveMoney().ServeHTTP).Methods("PUT")

	s.router.HandleFunc("/deal", handler.MakeDeal().ServeHTTP).Methods("PUT")
}

func (s *Server) constructConnString() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		s.config.Get("database.host"),
		s.config.Get("database.port"),
		s.config.Get("database.user"),
		s.config.Get("database.password"),
		s.config.Get("database.dbname"),
	)
}
