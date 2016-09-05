package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
)

func init() {
	SetSecret(uuid.NewV4().String())
}

func SetSecret(secretKey string) {
	secret = []byte(secretKey)
}

var secret []byte

func GenerateJwt(u *User) (*JWTAccessToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims(u))
	ss, err := token.SignedString(secret)

	var j JWTAccessToken
	j.TokenType = "Bearer"
	j.AccessToken = ss

	return &j, err
}

type JWTAccessToken struct {
	TokenType   string `json:"tokenType"`
	AccessToken string `json:"accessToken"`
}
