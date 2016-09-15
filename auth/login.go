package auth

import (
	"errors"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"net/http"
)

func LoginRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", LoginHandler)

	return r
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength <= 0 {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, errors.New("Request body is required"))
		return
	}

	u, err := bindUserFromRequest(r)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	if err := isValidLoginUser(u); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	if exists, err := u.Exists(); !exists || err != nil {
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err.Error())
			return
		}
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, "No Such User")
		return
	}

	if err := u.VerifyPassword(); err != nil {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, "Bad Password")
		return
	}

	ss, err := GenerateJwt(u)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
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
