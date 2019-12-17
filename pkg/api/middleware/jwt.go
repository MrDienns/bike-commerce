package middleware

import (
	"context"
	"crypto/rsa"
	"fmt"
	"net/http"
	"strings"

	"github.com/MrDienns/bike-commerce/pkg/api/response"
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	key *rsa.PublicKey
}

func NewJWT(key *rsa.PublicKey) *JWT {
	return &JWT{key}
}

func (m *JWT) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			response.WriteError(w, "Missing Authorization header", 401)
			return
		}
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return m.key, nil
		})
		if err != nil {
			response.WriteError(w, err.Error(), 401)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.WriteError(w, "Error mapping token", 401)
			return
		}

		if err = claims.Valid(); err != nil {
			response.WriteError(w, err.Error(), 401)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "session.user.id", claims["userid"])
		ctx = context.WithValue(ctx, "session.user.fullname", claims["name"])
		ctx = context.WithValue(ctx, "session.user.username", claims["username"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
