package controller

import (
	"crypto/rsa"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

type authRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type authResponse struct {
	Token string `json:"token"`
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

	var authRequest authRequest
	err = json.Unmarshal(body, &authRequest)
	if err != nil {
		response.WriteError(w, err.Error(), 422)
		return
	}

	crypto := sha1.New()
	crypto.Write([]byte(authRequest.Password))
	hash := crypto.Sum(nil)

	user := a.connector.UserFromCredentials(authRequest.Email, fmt.Sprintf("%x", hash))
	if user == nil {
		response.WriteError(w, "Onjuiste inloggegevens", 401)
		return
	}

	userBytes, _ := json.Marshal(user)
	var claims jwt.MapClaims
	json.Unmarshal(userBytes, &claims)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(a.privatekey)

	responseBytes, _ := json.Marshal(&authResponse{Token: tokenString})

	w.Write(responseBytes)
}
