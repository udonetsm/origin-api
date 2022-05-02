package controllers

import (
	"encoding/json"
	"net/http"
	"origin-api/getconf"

	"github.com/dgrijalva/jwt-go"
	"github.com/udonetsm/help/models"
)

func CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Auth")
		valid, err := TokenValid(token)
		if !valid {
			json.NewEncoder(w).Encode(models.ResponseAuth{Error: err.Error()})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func TokenValid(reqtoken string) (bool, error) {
	claims := models.Claims{}
	secret := getconf.Server.Secret
	token, err := jwt.ParseWithClaims(reqtoken, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	return token.Valid, err

}

func Test(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.ResponseAuth{Message: ":OK"})
}
