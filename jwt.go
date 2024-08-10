package main

import (
	"github.com/golang-jwt/jwt/v5"
)

func makeJwt(secret string, userId string) (string, error) {
	signingKey := []byte(secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:  "hw",
		Subject: userId,
	})

	return token.SignedString(signingKey)
}
