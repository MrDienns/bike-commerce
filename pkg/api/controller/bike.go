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

type Bike struct {
	logger   *zap.Logger
	key      *rsa.PublicKey
	bikeRepo database.Connector
}

func NewBike(logger *zap.Logger, key *rsa.PublicKey, bikeRepo database.Connector) *Bike {
	return &Bike{logger, key, bikeRepo}
}

func (b *Bike) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.NewJWT(b.key).Handle)

	r.Post("/", b.CreateBike)
	r.Get("/{id}", b.GetBike)
	r.Put("/{id}", b.UpdateBike)
	r.Delete("/{id}", b.DeleteBike)
	return r
}

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

func (b *Bike) DeleteBike(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := b.bikeRepo.DeleteBike(id)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	w.Write([]byte{})
}
