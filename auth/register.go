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
		render.JSON(w, http.StatusBadRequest, "Request body is required")
		return
	}
	u, err := bindUserFromRequest(r)
	if err != nil {
		render.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := isValidRegistrationUser(u); err != nil {
		render.JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if exists, err := u.Exists(); exists || err != nil {
		if err != nil {
			render.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}
		render.JSON(w, http.StatusConflict, "user exists")
		return
	}

	if err := u.HashPassword(); err != nil {
		render.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := u.Create(); err != nil {
		render.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	render.JSON(w, http.StatusCreated, u)
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
