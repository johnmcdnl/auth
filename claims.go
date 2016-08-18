package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
	"time"
)

var audience = "audience"
var issuer = "issuer"

func claims(u *User) *jwt.StandardClaims {
	var s jwt.StandardClaims

	now := time.Now()

	s.Id = uuid.NewV4().String()

	s.Audience = audience
	s.IssuedAt = now.Unix()
	s.ExpiresAt = now.Add(23 * time.Hour).Unix()

	s.Issuer = issuer

	s.Subject = u.Username

	return &s
}
