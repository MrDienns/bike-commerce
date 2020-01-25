package controller

import (
	"crypto/rsa"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MrDienns/bike-commerce/pkg/api/middleware"
	"github.com/MrDienns/bike-commerce/pkg/api/model"
	"github.com/MrDienns/bike-commerce/pkg/api/response"
	"github.com/MrDienns/bike-commerce/pkg/database"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type Accessory struct {
	logger        *zap.Logger
	key           *rsa.PublicKey
	accessoryRepo database.Connector
}

func NewAccessory(logger *zap.Logger, key *rsa.PublicKey, accessoryRepo database.Connector) *Accessory {
	return &Accessory{logger, key, accessoryRepo}
}

func (a *Accessory) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.NewJWT(a.key).Handle)

	r.Post("/", a.CreateAccessory)
	r.Get("/", a.GetAccessories)
	r.Get("/{id}", a.GetAccessory)
	r.Put("/{id}", a.UpdateAccessory)
	r.Delete("/{id}", a.DeleteAccessory)
	return r
}

func (a *Accessory) GetAccessories(w http.ResponseWriter, r *http.Request) {
	accessories, err := a.accessoryRepo.GetAccessories()
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	resp, _ := json.Marshal(accessories)
	w.Write(resp)
}

func (a *Accessory) GetAccessory(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	accessory, err := a.accessoryRepo.GetAccessory(id)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	resp, _ := json.Marshal(accessory)
	w.Write(resp)
}

func (a *Accessory) CreateAccessory(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var accessory model.Accessory
	decoder.Decode(&accessory)

	err := a.accessoryRepo.CreateAccessory(&accessory)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	w.Write([]byte{})
}

func (a *Accessory) UpdateAccessory(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	decoder := json.NewDecoder(r.Body)
	var accessory model.Accessory
	decoder.Decode(&accessory)

	accessory.ID = id

	err := a.accessoryRepo.UpdateAccessory(&accessory)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	w.Write([]byte{})
}

func (a *Accessory) DeleteAccessory(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := a.accessoryRepo.DeleteAccessory(id)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	w.Write([]byte{})
}
