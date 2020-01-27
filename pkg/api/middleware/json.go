package middleware

import (
	"net/http"

	"github.com/MrDienns/bike-commerce/pkg/api/response"
)

// Json is a middleware which ensures all content types are JSON.
type Json struct{}

// NewJson returns a new Json middleware
func NewJson() *Json {
	return &Json{}
}

// Handle overwrites the Handle method of an HTTP middleware.
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
