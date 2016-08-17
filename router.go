package main

import (
	"github.com/pressly/chi"
)

func AuthRouter() chi.Router {
	r := chi.NewRouter()

	r.Mount("/login", loginRouter())
	r.Mount("/register", registerRouter())
	r.Mount("/validate", validateRouter())

	return r
}


