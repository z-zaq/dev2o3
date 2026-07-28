package handlers

import (
	"net/http"

	"temux/internal/models"
	"temux/internal/repository"
	"temux/internal/utils"
	"temux/internal/services"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	TransactionRepo *repository.TransactionRepository
	WalletRepo      *repository.WalletRepository
	ReferralRepo    *repository.ReferralRepository
	WithdrawalRepo  *repository.WithdrawalRepository
	WalletService *services.WalletService
}

func (h *TransactionHandler) Deposit(
	c *gin.Context,
) {

	var req struct {
		Amount float64 `json:"amount"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	if req.Amount <= 0 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid amount"},
		)
		return
	}

	userID := utils.GetUserID(c)

	err := h.WalletService.Deposit(
	userID,
	req.Amount,
)

if err != nil {
	c.JSON(
		http.StatusInternalServerError,
		gin.H{
			"error": err.Error(),
		},
	)
	return
}
	//-----------------------------------
	// Referral Commission
	//-----------------------------------

	referrerID, err :=
		h.ReferralRepo.GetReferrerByReferredID(
			userID,
		)

	if err == nil {

		commission := req.Amount * 0.05

		_ = h.WalletRepo.AddBalance(
			referrerID,
			commission,
		)

		reward := &models.ReferralReward{
			ReferrerID:    referrerID,
			ReferredID:    userID,
			DepositAmount: req.Amount,
			Commission:    commission,
		}

		_ = h.ReferralRepo.CreateReward(
			reward,
		)
	}

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "deposit successful",
		},
	)
}
func (h *TransactionHandler) History(
	c *gin.Context,
) {

	userID := utils.GetUserID(c)

	transactions, err := h.TransactionRepo.GetByUserID(
		userID,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		transactions,
	)
}
func (h *TransactionHandler) Withdraw(
	c *gin.Context,
) {

	var req struct {
		Amount float64 `json:"amount"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	if req.Amount <= 0 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid amount"},
		)
		return
	}

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

	if wallet.Balance < req.Amount {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "insufficient balance"},
		)
		return
	}
	request := &models.Withdrawal{
		UserID: userID,
		Amount: req.Amount,
		Status: "pending",
	}

	err = h.WithdrawalRepo.Create(request)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "withdrawal request submitted",
		},
	)
}
