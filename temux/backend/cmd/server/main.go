package main

import (
	"log"

	"temux/internal/config"
	"temux/internal/database"
	"temux/internal/handlers"
	"temux/internal/middleware"
	"temux/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	//-----------------------------------
	// Load Environment Variables
	//-----------------------------------

	config.LoadEnv()

	//-----------------------------------
	// Initialize Database
	//-----------------------------------

	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	//-----------------------------------
	// Repositories
	//-----------------------------------

	userRepo := &repository.UserRepository{
		DB: db,
	}

	walletRepo := &repository.WalletRepository{
		DB: db,
	}

	//-----------------------------------
	// Handlers
	//-----------------------------------

	authHandler := &handlers.AuthHandler{
		Repo:       userRepo,
		WalletRepo: walletRepo,
	}

	walletHandler := &handlers.WalletHandler{
		WalletRepo: walletRepo,
	}

	//-----------------------------------
	// Router
	//-----------------------------------

	router := gin.Default()

	//-----------------------------------
	// Public Routes
	//-----------------------------------

	router.POST(
		"/api/register",
		authHandler.Register,
	)

	router.POST(
		"/api/login",
		authHandler.Login,
	)

	//-----------------------------------
	// Protected Routes
	//-----------------------------------

	api := router.Group("/api")

	api.Use(
		middleware.AuthMiddleware(),
	)

	api.GET(
		"/dashboard",
		func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome to Temux",
			})
		},
	)

	api.GET(
		"/wallet",
		walletHandler.GetWallet,
	)

	//-----------------------------------
	// Start Server
	//-----------------------------------

	log.Println("Server running on :8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}