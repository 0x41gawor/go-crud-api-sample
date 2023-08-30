package api

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(login string) (string, error) {
	claims := &jwt.MapClaims{
		"login":     login,
		"ExpiresAt": time.Now().Add(time.Minute * 15).Unix(),
	}

	secret := "SECRET"
	signingKey := []byte(secret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(signingKey)

	return tokenStr, err
}
