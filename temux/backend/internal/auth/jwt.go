package auth

import (
	"errors"
	"time"

	"temux/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(
	userID int,
) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().
			Add(24 * time.Hour).
			Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(config.JWTSecret()),
	)

}

func ParseToken(
	tokenString string,
) (jwt.MapClaims, error) {

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWTSecret()), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil

}
