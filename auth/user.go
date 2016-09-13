package auth

import (
	"encoding/json"
	"fmt"
	"github.com/pressly/chi/render"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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

func (u *User) HashPassword() error {
	h, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(h)
	return nil
}

func (u *User) VerifyPassword() error {

	conn, err := Connection()
	if err != nil {
		return err
	}
	c := conn.Get(u.Username)
	if err != nil {
		return err
	}
	r, err := c.Result()
	if err != nil {
		return err
	}
	var foundUser User
	if json.Unmarshal([]byte(r), &foundUser); err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(u.Password))
}

func (u *User) Create() error {
	userJson, err := json.Marshal(u)
	if err != nil {
		return err
	}
	conn, err := Connection()
	if err != nil {
		return err
	}
	c := conn.Set(u.Username, userJson, 0)
	if _, err := c.Result(); err != nil {
		return err
	}
	return nil
}

func (u *User) Exists() (bool, error) {
	conn, err := Connection()
	if err != nil {
		return false, err
	}
	c := conn.Get(u.Username)
	res, err := c.Result()
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
