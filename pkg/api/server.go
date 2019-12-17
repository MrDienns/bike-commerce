package api

import (
	"crypto/rsa"
	"net/http"

	"github.com/MrDienns/bike-commerce/pkg/api/middleware"

	"github.com/MrDienns/bike-commerce/pkg/api/controller"

	cm "github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type Server struct {
	logger *zap.Logger
	key    *rsa.PublicKey
}

func NewServer(logger *zap.Logger, key *rsa.PublicKey) *Server {
	return &Server{logger, key}
}

func (s *Server) Start() error {
	s.logger.Info("Starting HTTP server on port 8080")
	r := chi.NewRouter()

	r.Use(cm.Logger)
	r.Use(middleware.NewJson().Handle)

	r.Mount("/api", s.Routes())

	return http.ListenAndServe(":8080", r)
}

func (s *Server) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/user", controller.NewUser(s.logger, s.key).Routes())
	r.Mount("/authenticate", controller.NewAuthenticate(s.logger).Routes())
	return r
}
