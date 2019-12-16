package controller

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type Authenticate struct {
	logger *zap.Logger
}

func NewAuthenticate(logger *zap.Logger) *Authenticate {
	return &Authenticate{logger}
}

func (a *Authenticate) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", a.Login)
	return r
}

func (a *Authenticate) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("authenticate"))
}
