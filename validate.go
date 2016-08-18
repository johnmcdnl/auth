package main

import (
	"net/http"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"time"
	"strings"
)

func validateRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", validateHandler)

	return r
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")

	if err := validateJWT(auth); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	render.Status(r, http.StatusOK)
}

func validateJWT(s string) error {
	if s == "" {
		return errors.New("No Authorization header provided")
	}

	//TODO strip the bearer
	split := strings.Split(s, " ")

	if len(split) !=2 {
		return errors.New("Invalid token")
	}

	if split[0]!="Bearer"{
		return errors.New("Invalid Bearer token")
	}

	if err := verifyClaims(split[1]); err != nil {
		return err
	}

	return nil
}

func verifyClaims(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, keyFunc)

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	claims, claimsOk := token.Claims.(jwt.MapClaims)

	if !claimsOk {
		return errors.New("Claims not okay")
	}

	if !claims.VerifyIssuer(issuer, true) {
		return errors.New("VerifyIssuer")
	}

	if !claims.VerifyAudience(audience, true) {
		return errors.New("VerifyAudience")
	}

	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return errors.New("VerifyExpiresAt")
	}

	return nil
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return secret, nil
}