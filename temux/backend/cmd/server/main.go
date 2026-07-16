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
	transactionRepo := &repository.TransactionRepository{
		DB: db,
	}
	planRepo := &repository.PlanRepository{
		DB: db,
	}
	err = planRepo.SeedPlans()

	if err != nil {
		log.Fatal(err)
	}
	investmentRepo := &repository.InvestmentRepository{
		DB: db,
	}
	referralRepo := &repository.ReferralRepository{
		DB: db,
	}

	//-----------------------------------
	// Handlers
	//-----------------------------------

	authHandler := &handlers.AuthHandler{
		Repo:         userRepo,
		WalletRepo:   walletRepo,
		ReferralRepo: referralRepo,
	}

	walletHandler := &handlers.WalletHandler{
		WalletRepo: walletRepo,
	}
	transactionHandler := &handlers.TransactionHandler{
		TransactionRepo: transactionRepo,
		WalletRepo:      walletRepo,
		ReferralRepo:    referralRepo,
	}
	investmentHandler := &handlers.InvestmentHandler{
		InvestmentRepo: investmentRepo,
		PlanRepo:       planRepo,
		WalletRepo:     walletRepo,
	}
	dashboardHandler := &handlers.DashboardHandler{
		WalletRepo:      walletRepo,
		TransactionRepo: transactionRepo,
		InvestmentRepo:  investmentRepo,
	}
	referralHandler := &handlers.ReferralHandler{
		ReferralRepo: referralRepo,
	}
	adminHandler := &handlers.AdminHandler{
		UserRepo:        userRepo,
		TransactionRepo: transactionRepo,
		InvestmentRepo:  investmentRepo,
		ReferralRepo:    referralRepo,
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
	api.POST(
		"/deposit",
		transactionHandler.Deposit,
	)
	api.GET(
		"/transactions",
		transactionHandler.History,
	)
	api.POST(
		"/withdraw",
		transactionHandler.Withdraw,
	)
	api.POST(
		"/invest",
		investmentHandler.Invest,
	)
	api.GET(
		"/investments",
		investmentHandler.History,
	)
	api.GET(
		"/dashboard",
		dashboardHandler.GetDashboard,
	)
	api.GET(
		"/referrals",
		referralHandler.MyReferrals,
	)

	api.GET(
		"/referral-stats",
		referralHandler.Stats,
	)
	api.GET(
		"/referral-rewards",
		referralHandler.Rewards,
	)

	api.GET(
		"/referral-earnings",
		referralHandler.Earnings,
	)
	admin := api.Group("/admin")

	admin.Use(
		middleware.AdminMiddleware(
			userRepo,
		),
	)
	admin.GET(
		"/dashboard",
		adminHandler.Dashboard,
	)

	//-----------------------------------
	// Start Server
	//-----------------------------------

	log.Println("Server running on :8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
