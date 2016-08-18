package main

import (
	"net/http"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"errors"
)

func loginRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", loginHandler)

	return r
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength <= 0 {
		http.Error(w, "Request body is required", http.StatusBadRequest)
		return
	}

	u, err := bindUserFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := isValidLoginUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if exists, err := u.Exists(); !exists || err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Error(w, "No Such User", http.StatusUnauthorized)
		return
	}

	if err := u.VerifyPassword(); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	ss, err := GenerateJwt(u)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, nil)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, ss)
}

func isValidLoginUser(u *User) error {
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



