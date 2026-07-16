package handlers

import (
	"temux/internal/repository"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	UserRepo        *repository.UserRepository
	TransactionRepo *repository.TransactionRepository
	InvestmentRepo  *repository.InvestmentRepository
	ReferralRepo    *repository.ReferralRepository
}

func (h *AdminHandler) Dashboard(
	c *gin.Context,
) {

	totalUsers, _ :=
		h.UserRepo.CountUsers()

	totalDeposits, _ :=
		h.TransactionRepo.TotalDeposits()

	totalWithdrawals, _ :=
		h.TransactionRepo.TotalWithdrawals()

	totalInvestments, _ :=
		h.InvestmentRepo.TotalInvestments()

	activeInvestments, _ :=
		h.InvestmentRepo.TotalActiveInvestments()

	totalReferrals, _ :=
		h.ReferralRepo.TotalReferrals()

	c.JSON(
		200,
		gin.H{
			"total_users":        totalUsers,
			"total_deposits":     totalDeposits,
			"total_withdrawals":  totalWithdrawals,
			"total_investments":  totalInvestments,
			"active_investments": activeInvestments,
			"total_referrals":    totalReferrals,
		},
	)
}
