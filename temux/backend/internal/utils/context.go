package auth

import (
	"temux/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWTSecret()), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims, nil
}