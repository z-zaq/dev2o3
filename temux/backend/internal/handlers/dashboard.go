package handlers

import (
	"net/http"
	"temux/internal/models"
	"temux/internal/repository"
	"temux/internal/utils"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	WalletRepo      *repository.WalletRepository
	TransactionRepo *repository.TransactionRepository
	InvestmentRepo  *repository.InvestmentRepository
}

func (h *DashboardHandler) GetDashboard(
	c *gin.Context,
) {

	userID := utils.GetUserID(c)

	wallet, err := h.WalletRepo.GetWalletByUserID(
		userID,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	deposits, _ :=
		h.TransactionRepo.GetTotalDeposits(
			userID,
		)

	withdrawals, _ :=
		h.TransactionRepo.GetTotalWithdrawals(
			userID,
		)

	activeInvestments, _ :=
		h.InvestmentRepo.CountActive(
			userID,
		)

	totalInvested, _ :=
		h.InvestmentRepo.TotalInvested(
			userID,
		)

	totalProfit, _ :=
		h.InvestmentRepo.TotalProfit(
			userID,
		)

	data := models.Dashboard{
		WalletBalance:     wallet.Balance,
		TotalDeposits:     deposits,
		TotalWithdrawals:  withdrawals,
		ActiveInvestments: activeInvestments,
		TotalInvested:     totalInvested,
		TotalProfit:       totalProfit,
	}

	c.JSON(
		http.StatusOK,
		data,
	)
}
