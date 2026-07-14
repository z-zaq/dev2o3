package middleware

import (
	"net/http"
	"strings"

	"temux/internal/auth"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")

		if header == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "missing token"},
			)
			return
		}

		tokenString := strings.TrimPrefix(
			header,
			"Bearer ",
		)

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return auth.SecretKey, nil
			},
		)

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "invalid token"},
			)
			return
		}

		c.Next()
	}
}