package services

import (
	"temux/internal/repository"
)

type TransactionService struct {
	WalletService *WalletService
	ReferralRepo  *repository.ReferralRepository
}
