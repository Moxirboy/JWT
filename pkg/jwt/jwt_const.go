package json

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
