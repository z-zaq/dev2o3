package middleware

import (
	"net/http"
	"strings"
	"temux/internal/auth"

	"github.com/gin-gonic/gin"
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

		claims, err := auth.ParseToken(tokenString)

		if err != nil {
			// log.Println("ParseToken error:", err)
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "invalid token"},
			)
			return
		}

		userID := int(
			claims["user_id"].(float64),
		)

		c.Set("userID", userID)

		c.Next()
	}
}
