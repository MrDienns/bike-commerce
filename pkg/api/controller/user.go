package controller

import (
	"crypto/rsa"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MrDienns/bike-commerce/pkg/api/response"

	"github.com/MrDienns/bike-commerce/pkg/database"

	"github.com/MrDienns/bike-commerce/pkg/api/model"

	"github.com/MrDienns/bike-commerce/pkg/api/middleware"

	"go.uber.org/zap"

	"github.com/go-chi/chi"
)

type User struct {
	logger   *zap.Logger
	key      *rsa.PublicKey
	userRepo database.Connector
}

func NewUser(logger *zap.Logger, key *rsa.PublicKey, userRepo database.Connector) *User {
	return &User{logger, key, userRepo}
}

func (u *User) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.NewJWT(u.key).Handle)

	r.Post("/", u.CreateUser)
	r.Get("/", u.GetUsers)
	r.Get("/{id}", u.GetUser)
	r.Put("/{id}", u.UpdateUser)
	r.Delete("/{id}", u.DeleteUser)
	return r
}

func (u *User) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.userRepo.GetCustomers()
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	resp, _ := json.Marshal(users)
	w.Write(resp)
}

func (u *User) GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	user, err := u.userRepo.GetUser(id)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	resp, _ := json.Marshal(user)
	w.Write(resp)
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var user model.User
	decoder.Decode(&user)

	err := u.userRepo.CreateUser(&user)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	w.Write([]byte{})
}

func (u *User) UpdateUser(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	decoder := json.NewDecoder(r.Body)
	var user model.User
	decoder.Decode(&user)

	user.Id = id

	err := u.userRepo.UpdateUser(&user)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	w.Write([]byte{})
}

func (u *User) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := u.userRepo.DeleteUser(id)
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}
	w.Write([]byte{})
}
