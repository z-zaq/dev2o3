package main

import (
	"log"

	"temux/internal/database"
	"temux/internal/handlers"
	"temux/internal/middleware"
	"temux/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := database.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	repo := &repository.UserRepository{
		DB: db,
	}

	authHandler := &handlers.AuthHandler{
		Repo: repo,
	}

	router := gin.Default()

	router.POST(
		"/api/register",
		authHandler.Register,
	)

	router.POST(
		"/api/login",
		authHandler.Login,
	)

	protected := router.Group("/api")

	protected.Use(
		middleware.AuthMiddleware(),
	)

	protected.GET(
		"/dashboard",
		func(c *gin.Context) {
			c.JSON(200,
				gin.H{
					"message": "Welcome to Temux",
				})
		},
	)

	router.Run(":8080")
}