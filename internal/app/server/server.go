package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/srselivan/user-balance-microservice/internal/app/handler"
)

type Server struct {
	httpServer *http.Server
	router     *mux.Router
}

func (s *Server) Run(port string, handler *handler.Handler) error {
	s.router = mux.NewRouter()
	s.configureRouter(handler)

	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        s.router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	logrus.Infof("Starting server on %d port", port)

	return s.httpServer.ListenAndServe()
}

func (s *Server) configureRouter(handler *handler.Handler) {
	s.router.HandleFunc("/health", handler.Health().ServeHTTP).Methods("GET")

	s.router.HandleFunc("/balance", handler.GetBalance().ServeHTTP).Methods("GET")
	s.router.HandleFunc("/balance", handler.AppendBalance().ServeHTTP).Methods("PUT")

	s.router.HandleFunc("/holder", handler.FreezeAmount().ServeHTTP).Methods("POST")
	s.router.HandleFunc("/holder", handler.UnFreezeAmount().ServeHTTP).Methods("DELETE")
}
