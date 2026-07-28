package app

import (
	"database/sql"

	"temux/internal/repository"
	"temux/internal/services"
)

type Application struct {
	DB *sql.DB

	// Repositories
	UserRepo         *repository.UserRepository
	WalletRepo       *repository.WalletRepository
	TransactionRepo  *repository.TransactionRepository
	InvestmentRepo   *repository.InvestmentRepository
	PlanRepo         *repository.PlanRepository
	ReferralRepo     *repository.ReferralRepository
	WithdrawalRepo   *repository.WithdrawalRepository
	PaymentRepo      *repository.PaymentRepository
	AuditRepo        *repository.AuditRepository

	// Services
	WalletService      *services.WalletService
	TransactionService *services.TransactionService
	AuditService       *services.AuditService

}
func NewApplication(
	db *sql.DB,
) *Application {

	app := &Application{
		DB: db,
	}

	//-----------------------------------
	// Repositories
	//-----------------------------------

	app.UserRepo = &repository.UserRepository{
		DB: db,
	}

	app.WalletRepo = &repository.WalletRepository{
		DB: db,
	}

	app.TransactionRepo = &repository.TransactionRepository{
		DB: db,
	}

	app.PlanRepo = &repository.PlanRepository{
		DB: db,
	}

	app.InvestmentRepo = &repository.InvestmentRepository{
		DB: db,
	}

	app.ReferralRepo = &repository.ReferralRepository{
		DB: db,
	}

	app.WithdrawalRepo = &repository.WithdrawalRepository{
		DB: db,
	}

	app.PaymentRepo = &repository.PaymentRepository{
		DB: db,
	}

	app.AuditRepo = &repository.AuditRepository{
		DB: db,
	}

	//-----------------------------------
	// Services
	//-----------------------------------

	app.WalletService = &services.WalletService{
		DB:              db,
		WalletRepo:      app.WalletRepo,
		TransactionRepo: app.TransactionRepo,
	}

	app.TransactionService = &services.TransactionService{
		WalletService: app.WalletService,
		ReferralRepo:  app.ReferralRepo,
	}

	app.AuditService = &services.AuditService{
		AuditRepo: app.AuditRepo,
	}

	return app
}