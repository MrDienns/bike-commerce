package controller

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/go-chi/chi"
)

type User struct {
	logger *zap.Logger
}

func NewUser(logger *zap.Logger) *User {
	return &User{logger}
}

func (u *User) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", u.CreateUser)
	r.Get("/{id}", u.GetUser)
	r.Put("/{id}", u.UpdateUser)
	r.Delete("/{id}", u.DeleteUser)
	return r
}

func (u *User) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create user"))
}

func (u *User) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update user"))
}

func (u *User) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user"))
}
