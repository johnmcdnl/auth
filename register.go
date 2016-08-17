package main

import (
	"net/http"
	"github.com/pressly/chi"
	"errors"
	"github.com/pressly/chi/render"
)

func registerRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", registerHandler)

	return r
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	u, err := bindUserFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := isValidRegistrationUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if exists, err := u.Exists(); exists || err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Error(w, "user exists", http.StatusConflict)
		return
	}

	if err := u.HashPassword(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := u.Create(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, u)
}

func isValidRegistrationUser(u *User) error {
	if u == nil {
		return errors.New("nil user")
	}
	if u.Username == "" {
		return errors.New("nil Username")
	}
	if u.Password == "" {
		return errors.New("nil Password")
	}

	return nil
}