package controller

import (
	"crypto/rsa"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MrDienns/bike-commerce/pkg/api/middleware"
	"github.com/MrDienns/bike-commerce/pkg/api/model"
	"github.com/MrDienns/bike-commerce/pkg/api/response"
	"github.com/go-chi/chi"

	"github.com/MrDienns/bike-commerce/pkg/database"
	"go.uber.org/zap"
)

// Bike controller is a controller which accepts bike related requests.
type Bike struct {
	logger   *zap.Logger
	key      *rsa.PublicKey
	bikeRepo database.Connector
}

// NewBike creates a new Bike controller with a set of parameters and return sit.
func NewBike(logger *zap.Logger, key *rsa.PublicKey, bikeRepo database.Connector) *Bike {
	return &Bike{logger, key, bikeRepo}
}

// Routes returns a *chi.Mux object with all endpoints registered on it.
func (b *Bike) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.NewJWT(b.key).Handle)

	r.Post("/", b.CreateBike)
	r.Get("/", b.GetBikes)
	r.Get("/{id}", b.GetBike)
	r.Put("/{id}", b.UpdateBike)
	r.Delete("/{id}", b.DeleteBike)
	return r
}

// GetBikes returns a list of all bikes.
func (b *Bike) GetBikes(w http.ResponseWriter, r *http.Request) {
	bikes, err := b.bikeRepo.GetBikes()
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	resp, _ := json.Marshal(bikes)
	w.Write(resp)
}

// GetBike returns a bike based on the provided ID.
func (b *Bike) GetBike(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	bike, err := b.bikeRepo.GetBike(id)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	resp, _ := json.Marshal(bike)
	w.Write(resp)
}

// CreateBike accepts a bike as body and creates it.
func (b *Bike) CreateBike(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var bike model.Bike
	decoder.Decode(&bike)

	err := b.bikeRepo.CreateBike(&bike)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	w.Write([]byte{})
}

// UpdateBike accepts a bike as body and updates an existing one.
func (b *Bike) UpdateBike(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	decoder := json.NewDecoder(r.Body)
	var bike model.Bike
	decoder.Decode(&bike)

	bike.ID = id

	err := b.bikeRepo.UpdateBike(&bike)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	w.Write([]byte{})
}

// DeleteBike accepts a bike ID and deletes it.
func (b *Bike) DeleteBike(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := b.bikeRepo.DeleteBike(id)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	w.Write([]byte{})
}
