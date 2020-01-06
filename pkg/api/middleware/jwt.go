package middleware

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/MrDienns/bike-commerce/pkg/api/model"

	"github.com/MrDienns/bike-commerce/pkg/api/response"
	"github.com/dgrijalva/jwt-go"
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

		user, err := m.userFromToken(tokenStr)
		if err != nil {
			response.WriteError(w, err.Error(), 401)
			return
		}

		ctx := r.Context()

		ctx = context.WithValue(ctx, "session.user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// userFromToken takes the token string and tries to transform it into a *model.User object.
func (m *JWT) userFromToken(tokenStr string) (*model.User, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return m.key, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	tokenJson, _ := json.Marshal(claims)

	var user model.User
	err = json.Unmarshal(tokenJson, &user)
	if err != nil {
		return nil, err
	}

	if err = claims.Valid(); err != nil {
		return nil, err
	}

	return &user, nil
}
