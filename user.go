package main

import (
	"github.com/pressly/chi/render"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"encoding"
)

type User struct {
	Username string
	Password string
}

func bindUserFromRequest(r *http.Request) (*User, error) {
	var u *User

	if err := render.Bind(r.Body, &u); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User)HashPassword() error {
	h, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(h)
	return nil
}

func (u *User)Create() error {
	cmd := Connection().Set(u.Username, encoding.BinaryMarshaler(u), 0)
	if _, err := cmd.Result(); err != nil {
		return err
	}
	return nil
}

func (u *User)Exists() (bool, error) {
	cmd := Connection().Get(u.Username)

	res, err := cmd.Result()
	if res == "" {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *User) MarshalBinary() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", u)), nil
}
