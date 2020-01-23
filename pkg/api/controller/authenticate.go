package controller

import (
	"crypto/rsa"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/MrDienns/bike-commerce/pkg/api/model"

	"github.com/dgrijalva/jwt-go"

	"github.com/MrDienns/bike-commerce/pkg/api/response"

	"github.com/MrDienns/bike-commerce/pkg/database"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type Authenticate struct {
	logger     *zap.Logger
	publickey  *rsa.PublicKey
	privatekey *rsa.PrivateKey
	connector  database.Connector
}

func NewAuthenticate(logger *zap.Logger, publickey *rsa.PublicKey, privatekey *rsa.PrivateKey,
	connector database.Connector) *Authenticate {
	return &Authenticate{logger, publickey, privatekey, connector}
}

func (a *Authenticate) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", a.Login)
	return r
}

func (a *Authenticate) Login(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.WriteError(w, err.Error(), 500)
		return
	}

	var authRequest model.AuthRequest
	err = json.Unmarshal(body, &authRequest)
	if err != nil {
		response.WriteError(w, err.Error(), 422)
		return
	}

	user, err := a.connector.UserFromCredentials(authRequest.Email, authRequest.Password)
	if err != nil {
		response.WriteError(w, err.Error(), 401)
		return
	}
	if user == nil {
		response.WriteError(w, "Onjuiste inloggegevens", 401)
		return
	}

	userBytes, _ := json.Marshal(user)
	var claims jwt.MapClaims
	json.Unmarshal(userBytes, &claims)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(a.privatekey)

	responseBytes, _ := json.Marshal(&response.AuthResponse{Token: tokenString})

	w.Write(responseBytes)
}
