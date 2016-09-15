package auth

import (
	"errors"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"net/http"
)

func RegisterRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", RegisterHandler)

	return r
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "Request body is required")
		return
	}
	u, err := bindUserFromRequest(r)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}

	if err := isValidRegistrationUser(u); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	if exists, err := u.Exists(); exists || err != nil {
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err.Error())
			return
		}
		render.Status(r, http.StatusConflict)
		render.JSON(w, r, "user exists")
		return
	}

	if err := u.HashPassword(); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}

	if err := u.Create(); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
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
