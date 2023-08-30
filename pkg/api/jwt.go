package api

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(login string) (string, error) {
	claims := &jwt.MapClaims{
		"login":     login,
		"expiresAt": time.Now().Add(time.Minute * 15).Unix(),
	}

	secret := "SECRET"
	signingKey := []byte(secret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(signingKey)

	return tokenStr, err
}

func ValidateJWT(tokenStr string) (*jwt.Token, error) {
	secret := "SECRET"

	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}
