package main

import (
	"net/http"
	"github.com/pressly/chi"
)

func validateRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", validateHandler)

	return r
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("validateHandler"))
}
