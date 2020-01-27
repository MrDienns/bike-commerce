package controller

import (
	"crypto/rsa"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MrDienns/bike-commerce/pkg/api/response"

	"github.com/MrDienns/bike-commerce/pkg/api/middleware"

	"github.com/MrDienns/bike-commerce/pkg/api/model"

	"github.com/go-chi/chi"

	"github.com/MrDienns/bike-commerce/pkg/database"
	"go.uber.org/zap"
)

// Customer controller is a controller which accepts customer related requests.
type Customer struct {
	logger       *zap.Logger
	key          *rsa.PublicKey
	customerRepo database.Connector
}

// NewCustomer creates the customer controller and returns it.
func NewCustomer(logger *zap.Logger, key *rsa.PublicKey, customerRepo database.Connector) *Customer {
	return &Customer{logger, key, customerRepo}
}

// Routes returns a *chi.Mux with all endpoints registered on it.
func (c *Customer) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.NewJWT(c.key).Handle)

	r.Post("/", c.CreateCustomer)
	r.Get("/", c.GetCustomers)
	r.Get("/{id}", c.GetCustomer)
	r.Put("/{id}", c.UpdateCustomer)
	r.Delete("/{id}", c.DeleteCustomer)
	return r
}

// GetCustomers returns all customers.
func (c *Customer) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := c.customerRepo.GetCustomers()
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	resp, _ := json.Marshal(customers)
	w.Write(resp)
}

// GetCustomer accepts an ID an returns the appropriate customer.
func (c *Customer) GetCustomer(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	customer, err := c.customerRepo.GetCustomer(id)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	resp, _ := json.Marshal(customer)
	w.Write(resp)
}

// CreateCustomer accepts a customer model and creates it.
func (c *Customer) CreateCustomer(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var customer model.Customer
	decoder.Decode(&customer)

	err := c.customerRepo.CreateCustomer(&customer)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	w.Write([]byte{})
}

// UpdateCustomer accepts a customer object and updates an existing one.
func (c *Customer) UpdateCustomer(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	decoder := json.NewDecoder(r.Body)
	var customer model.Customer
	decoder.Decode(&customer)

	customer.ID = id

	err := c.customerRepo.UpdateCustomer(&customer)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	w.Write([]byte{})
}

// DeleteCustomer accepts an ID and deletes the customer.
func (c *Customer) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := c.customerRepo.DeleteCustomer(id)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	w.Write([]byte{})
}
