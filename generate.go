package main

import (
	"github.com/dgrijalva/jwt-go"
	//"github.com/satori/go.uuid"
	"fmt"
)

func init(){
	fmt.Println(string(secret))
}

//var secret = []byte(uuid.NewV4().String())
var secret = []byte("247af323-6c90-4b7b-b836-ef2cccb43c6d")

func GenerateJwt(u *User) (*JWTAccessToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims(u))
	ss, err := token.SignedString(secret)

	var j JWTAccessToken
	j.TokenType = "Bearer"
	j.AccessToken = ss

	return &j, err
}

type JWTAccessToken struct {
	TokenType string `json:"tokenType"`
	AccessToken     string `json:"accessToken"`
}