package middleware

import (
	"context"
	"crypto/rsa"
	"net/http"
	"strings"

	"github.com/MrDienns/bike-commerce/pkg/util"

	"github.com/MrDienns/bike-commerce/pkg/api/response"
)

// JWT is a struct in the controller package, which acts as an authentication middleware by parsing the JWT from the
// request and setting it on the request context. If no JWT token was sent, an HTTP 401 is returned.
type JWT struct {
	key *rsa.PublicKey
}

// NewJWT accepts an *rsa.PublicKey and returns new *JWT instance.
func NewJWT(key *rsa.PublicKey) *JWT {
	return &JWT{key}
}

// Handle is the middleware implementation function that gets called when requests come in
func (m *JWT) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			response.WriteError(w, "Missing Authorization header", 401)
			return
		}
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		user, err := util.UserFromToken(m.key, tokenStr)
		if err != nil {
			response.WriteError(w, err.Error(), 401)
			return
		}

		ctx := r.Context()

		ctx = context.WithValue(ctx, "session.user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
