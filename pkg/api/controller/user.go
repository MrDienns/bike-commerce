package controller

import (
	"crypto/rsa"
	"net/http"

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
	r.Get("/{id}", u.GetUser)
	r.Put("/{id}", u.UpdateUser)
	r.Delete("/{id}", u.DeleteUser)
	return r
}

func (u *User) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Context().Value("session.user").(*model.User).Name))
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Context().Value("session.user").(*model.User).Name))
}

func (u *User) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Context().Value("session.user").(*model.User).Name))
}

func (u *User) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Context().Value("session.user").(*model.User).Name))
}
