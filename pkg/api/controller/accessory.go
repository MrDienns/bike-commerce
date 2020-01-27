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

// Accessory is a controller which handles the REST API endpoints for accessories.
type Accessory struct {
	logger        *zap.Logger
	key           *rsa.PublicKey
	accessoryRepo database.Connector
}

// NewAccessory creates a new accessory controller.
func NewAccessory(logger *zap.Logger, key *rsa.PublicKey, accessoryRepo database.Connector) *Accessory {
	return &Accessory{logger, key, accessoryRepo}
}

// Routes returns a *chi.Mux routes with all endpoints added to it.
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

// GetAccessories returns a list of accessories.
func (a *Accessory) GetAccessories(w http.ResponseWriter, r *http.Request) {
	accessories, err := a.accessoryRepo.GetAccessories()
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	resp, _ := json.Marshal(accessories)
	w.Write(resp)
}

// GetAccessory returns the accessory based on the provided ID.
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

// CreateAccessory accepts an accessory as body and creates it.
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

// UpdateAccessory takes an accessory as body and updates an existing one.
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

// DeleteAccessory takes an accessory ID and deletes it.
func (a *Accessory) DeleteAccessory(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := a.accessoryRepo.DeleteAccessory(id)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	w.Write([]byte{})
}
