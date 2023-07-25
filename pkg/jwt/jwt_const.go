package json

import (
	"github.com/caarlos0/env/v6"
	"github.com/dgrijalva/jwt-go"
)

type jWT struct {
	signingKey string `env:"SIGNING"`
	salt       string `env:"SALT"`
	tokenTTL   string `env:"TOKENTTL"`
}

var Jwt jWT

func GetterSigningKey() *jWT {
	if err := env.Parse(&Jwt); err != nil {
		panic(err)
	}
	return &Jwt
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
