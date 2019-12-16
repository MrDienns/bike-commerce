package api

import (
	"net/http"

	"github.com/MrDienns/bike-commerce/pkg/api/controller"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type Server struct {
	logger *zap.Logger
}

func NewServer(logger *zap.Logger) *Server {
	return &Server{logger}
}

func (s *Server) Start() error {
	s.logger.Info("Starting HTTP server on port 8080")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/api", s.Routes())

	return http.ListenAndServe(":8080", r)
}

func (s *Server) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/user", controller.NewUser(s.logger).Routes())
	r.Mount("/authenticate", controller.NewAuthenticate(s.logger).Routes())
	return r
}
