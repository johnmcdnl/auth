package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
	"time"
)

var audience = "audience"

type CustomClaims struct {
	*jwt.StandardClaims
}

func claims(u *User) CustomClaims {
	var c CustomClaims
	var s jwt.StandardClaims

	now := time.Now()

	s.Id = uuid.NewV4().String()

	s.Audience = audience
	s.IssuedAt = now.Unix()
	s.NotBefore = now.Unix()
	s.ExpiresAt = now.Add(100 * time.Minute).Unix()

	s.Issuer = "Issuer"

	s.Subject = u.Username

	c.StandardClaims = &s

	return c
}
