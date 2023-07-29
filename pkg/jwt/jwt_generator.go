package json

import (
	repository "JWT/internal/service/repo"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

import (
	"errors"
)

func GenerateToken(username string, password string) (string, error) {
	user, err := repository.GetUser(username, GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(Jwt.signingKey))
}
func ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(Jwt.signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not type of*tokenClaims ")
	}
	return claims.UserId, nil
}
func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	Jwt.salt = "12345"
	return fmt.Sprintf("%x", hash.Sum([]byte(Jwt.salt)))
}
