package utils

import (
	"temux/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetUserID(c *gin.Context) int {
	userID, exists := c.Get("userID")

	if !exists {
		return 0
	}

	return userID.(int)
}
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
