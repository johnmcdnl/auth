package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

func ValidateHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		aHeader := r.Header.Get("Authorization")

		if err := validateJWT(aHeader); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}


func validateJWT(s string) error {
	if s == "" {
		return errors.New("No Authorization header provided")
	}

	split := strings.Split(s, " ")

	if len(split) != 2 {
		return errors.New("Invalid token")
	}

	if split[0] != "Bearer" {
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