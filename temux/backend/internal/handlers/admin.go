package handlers

import (
	"strconv"
	"temux/internal/models"
	"temux/internal/repository"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	UserRepo        *repository.UserRepository
	WalletRepo      *repository.WalletRepository
	TransactionRepo *repository.TransactionRepository
	InvestmentRepo  *repository.InvestmentRepository
	ReferralRepo    *repository.ReferralRepository
	WithdrawalRepo  *repository.WithdrawalRepository
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
func (h *AdminHandler) Users(
	c *gin.Context,
) {

	users, err :=
		h.UserRepo.GetAllUsers()

	if err != nil {

		c.JSON(
			500,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		200,
		users,
	)
}
func (h *AdminHandler) UserDetails(
	c *gin.Context,
) {

	idParam := c.Param("id")

	userID, err := strconv.Atoi(idParam)

	if err != nil {

		c.JSON(
			400,
			gin.H{
				"error": "invalid user id",
			},
		)

		return
	}

	user, err :=
		h.UserRepo.GetByID(userID)

	if err != nil {

		c.JSON(
			404,
			gin.H{
				"error": "user not found",
			},
		)

		return
	}

	wallet, _ :=
		h.WalletRepo.GetWalletByUserID(
			userID,
		)

	totalDeposits, _ :=
		h.TransactionRepo.GetTotalDeposits(
			userID,
		)

	totalWithdrawals, _ :=
		h.TransactionRepo.GetTotalWithdrawals(
			userID,
		)

	activeInvestments, _ :=
		h.InvestmentRepo.CountActiveByUser(
			userID,
		)

	referrals, _ :=
		h.ReferralRepo.CountReferrals(
			userID,
		)

	c.JSON(
		200,
		gin.H{
			"user":               user,
			"wallet":             wallet,
			"total_deposits":     totalDeposits,
			"total_withdrawals":  totalWithdrawals,
			"active_investments": activeInvestments,
			"referral_count":     referrals,
		},
	)
}
func (h *AdminHandler) Transactions(
	c *gin.Context,
) {

	transactions, err :=
		h.TransactionRepo.GetAllTransactions()

	if err != nil {

		c.JSON(
			500,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		200,
		transactions,
	)
}
func (h *AdminHandler) PendingWithdrawals(
	c *gin.Context,
) {

	withdrawals, err :=
		h.WithdrawalRepo.GetPending()

	if err != nil {

		c.JSON(
			500,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		200,
		withdrawals,
	)
}
func (h *AdminHandler) ApproveWithdrawal(
	c *gin.Context,
) {

	idParam := c.Param("id")

	withdrawalID, err := strconv.Atoi(
		idParam,
	)

	if err != nil {
		c.JSON(
			400,
			gin.H{
				"error": "invalid withdrawal id",
			},
		)
		return
	}

	withdrawal, err :=
		h.WithdrawalRepo.GetByID(
			withdrawalID,
		)

	if err != nil {
		c.JSON(
			404,
			gin.H{
				"error": "withdrawal not found",
			},
		)
		return
	}

	if withdrawal.Status != "pending" {
		c.JSON(
			400,
			gin.H{
				"error": "withdrawal already processed",
			},
		)
		return
	}

	wallet, err :=
		h.WalletRepo.GetWalletByUserID(
			withdrawal.UserID,
		)

	if err != nil {
		c.JSON(
			500,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	if wallet.Balance < withdrawal.Amount {
		c.JSON(
			400,
			gin.H{
				"error": "insufficient wallet balance",
			},
		)
		return
	}

	err = h.WalletRepo.DeductBalance(
		withdrawal.UserID,
		withdrawal.Amount,
	)

	if err != nil {
		c.JSON(
			500,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	tx := &models.Transaction{
		UserID: withdrawal.UserID,
		Type:   "withdraw",
		Amount: withdrawal.Amount,
	}

	err = h.TransactionRepo.CreateTransaction(
		tx,
	)

	if err != nil {
		c.JSON(
			500,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	err = h.WithdrawalRepo.UpdateStatus(
		withdrawal.ID,
		"approved",
	)

	if err != nil {
		c.JSON(
			500,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		200,
		gin.H{
			"message": "withdrawal approved",
		},
	)
}

// func (h *AdminHandler) ApproveWithdrawal(
// 	c *gin.Context,
// ) {

//		c.JSON(
//			200,
//			gin.H{
//				"message": "approve withdrawal not implemented yet",
//			},
//		)
//	}
func (h *AdminHandler) RejectWithdrawal(
	c *gin.Context,
) {

	c.JSON(
		200,
		gin.H{
			"message": "withdrawal rejected",
		},
	)
}
