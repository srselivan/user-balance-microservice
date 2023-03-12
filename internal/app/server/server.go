package server

import (
	"context"
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

	logrus.Infof("Starting server on %v port", port)

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) configureRouter(handler *handler.Handler) {
	s.router.HandleFunc("/api/v1/health", handler.Health().ServeHTTP).Methods("GET")

	s.router.HandleFunc("/api/v1/balance", handler.GetBalance().ServeHTTP).Methods("GET")
	s.router.HandleFunc("/api/v1/balance", handler.AppendBalance().ServeHTTP).Methods("PUT")

	s.router.HandleFunc("/api/v1/holder", handler.FreezeAmount().ServeHTTP).Methods("POST")
	s.router.HandleFunc("/api/v1/holder", handler.UnFreezeAmount().ServeHTTP).Methods("DELETE")
}
