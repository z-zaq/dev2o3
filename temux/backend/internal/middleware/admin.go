package middleware

import (
	"net/http"

	"temux/internal/repository"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware(
	userRepo *repository.UserRepository,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		userIDAny, exists := c.Get("userID")

		if !exists {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "unauthorized"},
			)
			return
		}

		userID := userIDAny.(int)

		user, err := userRepo.GetByID(userID)

		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "user not found"},
			)
			return
		}

		if !user.IsAdmin {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{"error": "admin only"},
			)
			return
		}

		c.Next()
	}
}
