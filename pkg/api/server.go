package api

import (
	"crypto/rsa"
	"net/http"

	"github.com/MrDienns/bike-commerce/pkg/database"

	"github.com/MrDienns/bike-commerce/pkg/api/middleware"

	"github.com/MrDienns/bike-commerce/pkg/api/controller"

	cm "github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

// Server is the HTTP API server.
type Server struct {
	logger     *zap.Logger
	publickey  *rsa.PublicKey
	privatekey *rsa.PrivateKey
	connector  database.Connector
}

// NewServer creates and returns a new HTTP server.
func NewServer(logger *zap.Logger, publickey *rsa.PublicKey, privatekey *rsa.PrivateKey,
	connector database.Connector) *Server {
	return &Server{logger, publickey, privatekey, connector}
}

// Start starts the HTTP server.
func (s *Server) Start() error {
	s.logger.Info("Starting HTTP server on port 8080")
	r := chi.NewRouter()

	r.Use(cm.Logger)
	r.Use(middleware.NewJson().Handle)

	r.Mount("/api", s.Routes())

	return http.ListenAndServe(":8080", r)
}

// Routes returns a *chi.Mux object with all endpoints registered on it.
func (s *Server) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/rental", controller.NewRental(s.logger, s.publickey, s.connector).Routes())
	r.Mount("/accessory", controller.NewAccessory(s.logger, s.publickey, s.connector).Routes())
	r.Mount("/bike", controller.NewBike(s.logger, s.publickey, s.connector).Routes())
	r.Mount("/user", controller.NewUser(s.logger, s.publickey, s.connector).Routes())
	r.Mount("/customer", controller.NewCustomer(s.logger, s.publickey, s.connector).Routes())
	r.Mount("/authenticate", controller.NewAuthenticate(s.logger, s.publickey, s.privatekey, s.connector).Routes())
	return r
}
