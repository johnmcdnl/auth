package main

import (
	"net/http"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
)

func loginRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", loginHandler)

	return r
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	var u *User

	if err := render.Bind(r.Body, &u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !isValidUser(u) {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, nil)
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

func isValidUser(u *User) bool {
	//TODO
	return true
}