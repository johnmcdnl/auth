package auth

import (
	"github.com/pressly/chi"
)

func AuthRouter() chi.Router {
	r := chi.NewRouter()

	r.Mount("/login", LoginRouter())
	r.Mount("/register", RegisterRouter())

	return r
}