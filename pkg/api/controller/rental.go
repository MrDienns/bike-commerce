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

type Rental struct {
	logger     *zap.Logger
	key        *rsa.PublicKey
	rentalRepo database.Connector
}

func NewRental(logger *zap.Logger, key *rsa.PublicKey, rentalRepo database.Connector) *Rental {
	return &Rental{logger, key, rentalRepo}
}

func (r *Rental) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.NewJWT(r.key).Handle)

	router.Post("/", r.CreateRental)
	router.Get("/", r.GetRentals)
	router.Get("/{id}", r.GetRental)
	router.Put("/{id}", r.UpdateRental)
	router.Delete("/{id}", r.DeleteRental)
	return router
}

func (r *Rental) GetRentals(w http.ResponseWriter, req *http.Request) {
	rentals, err := r.rentalRepo.GetRentals()
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	resp, _ := json.Marshal(rentals)
	w.Write(resp)
}

func (r *Rental) GetRental(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(req, "id"))
	rental, err := r.rentalRepo.GetRental(id)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	resp, _ := json.Marshal(rental)
	w.Write(resp)
}

func (r *Rental) CreateRental(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var rental model.Rental
	decoder.Decode(&rental)

	err := r.rentalRepo.CreateRental(&rental)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	w.Write([]byte{})
}

func (r *Rental) UpdateRental(w http.ResponseWriter, req *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(req, "id"))

	decoder := json.NewDecoder(req.Body)
	var rental model.Rental
	decoder.Decode(&rental)

	rental.ID = id

	err := r.rentalRepo.UpdateRental(&rental)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	w.Write([]byte{})
}

func (r *Rental) DeleteRental(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(req, "id"))
	err := r.rentalRepo.DeleteRental(id)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	w.Write([]byte{})
}
