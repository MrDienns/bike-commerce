package util

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"

	"github.com/MrDienns/bike-commerce/pkg/api/model"
	"github.com/dgrijalva/jwt-go"
)

// UserFromToken takes the public key and the token string and tries to transform it into a *model.User object.
func UserFromToken(key *rsa.PublicKey, tokenStr string) (*model.User, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
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
