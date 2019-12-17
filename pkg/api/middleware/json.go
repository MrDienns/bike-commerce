package middleware

import (
	"net/http"

	"github.com/MrDienns/bike-commerce/pkg/api/response"
)

type Json struct{}

func NewJson() *Json {
	return &Json{}
}

func (j *Json) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Header.Get("Content-Type") != "application/json" {
			response.WriteError(w, "Invalid Content-Type", 415)
			return
		}
		next.ServeHTTP(w, r)
	})
}
