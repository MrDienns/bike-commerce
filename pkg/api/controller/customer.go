package controller

import (
	"crypto/rsa"
	"net/http"

	"github.com/MrDienns/bike-commerce/pkg/api/middleware"

	"github.com/MrDienns/bike-commerce/pkg/api/model"

	"github.com/go-chi/chi"

	"github.com/MrDienns/bike-commerce/pkg/database"
	"go.uber.org/zap"
)

type Customer struct {
	logger       *zap.Logger
	key          *rsa.PublicKey
	customerRepo database.Connector
}

func NewCustomer(logger *zap.Logger, key *rsa.PublicKey, customerRepo database.Connector) *Customer {
	return &Customer{logger, key, customerRepo}
}

func (c *Customer) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.NewJWT(c.key).Handle)

	r.Post("/", c.CreateCustomer)
	r.Get("/{id}", c.GetCustomer)
	r.Put("/{id}", c.UpdateCustomer)
	r.Delete("/{id}", c.DeleteCustomer)
	return r
}

func (c *Customer) GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Context().Value("session.user").(*model.User).Name))
}

func (c *Customer) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Context().Value("session.user").(*model.User).Name))
}

func (c *Customer) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Context().Value("session.user").(*model.User).Name))
}

func (c *Customer) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Context().Value("session.user").(*model.User).Name))
}