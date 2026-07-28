package main

import (
	"log"

	"temux/internal/config"
	"temux/internal/database"
	"temux/internal/handlers"
	"temux/internal/middleware"
	"temux/internal/repository"
	"temux/internal/services"
	"temux/internal/scheduler"
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
	withdrawalRepo := &repository.WithdrawalRepository{
		DB: db,
	}
	paymentHandler := &handlers.PaymentHandler{
	UserRepo: userRepo,
}
walletService := &services.WalletService{
	DB:              db,
	WalletRepo:      walletRepo,
	TransactionRepo: transactionRepo,
}
sch := scheduler.New()

sch.Register(&scheduler.ProfitJob{
	InvestmentRepo: investmentRepo,
})

sch.Start()

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
		WithdrawalRepo:  withdrawalRepo,
		WalletService:   walletService,
	}
	investmentHandler := &handlers.InvestmentHandler{
		InvestmentRepo:  investmentRepo,
		PlanRepo:        planRepo,
		WalletRepo:      walletRepo,
		TransactionRepo: transactionRepo,
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
		WalletRepo:      walletRepo,
		TransactionRepo: transactionRepo,
		InvestmentRepo:  investmentRepo,
		ReferralRepo:    referralRepo,
		WithdrawalRepo:  withdrawalRepo,
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
	//-----------------------------------
	// Admin Routes
	//-----------------------------------

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
	admin.GET(
		"/users",
		adminHandler.Users,
	)
	admin.GET(
		"/users/:id",
		adminHandler.UserDetails,
	)
	admin.GET(
		"/transactions",
		adminHandler.Transactions,
	)
	admin.GET(
		"/withdrawals",
		adminHandler.PendingWithdrawals,
	)
	admin.POST(
		"/withdrawals/:id/approve",
		adminHandler.ApproveWithdrawal,
	)

	admin.POST(
		"/withdrawals/:id/reject",
		adminHandler.RejectWithdrawal,
	)
	api.POST(
	"/payments/deposit",
	paymentHandler.InitializeDeposit,
)

	//-----------------------------------
	// Start Server
	//-----------------------------------

	log.Println("Server running on :8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
