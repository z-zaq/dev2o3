package handlers

import (
	"net/http"
	"temux/internal/models"
	"temux/internal/repository"
	"temux/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type InvestmentHandler struct {
	InvestmentRepo *repository.InvestmentRepository
	PlanRepo       *repository.PlanRepository
	WalletRepo     *repository.WalletRepository
}

func (h *InvestmentHandler) Invest(
	c *gin.Context,
) {

	var req struct {
		PlanID int     `json:"plan_id"`
		Amount float64 `json:"amount"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	userID := utils.GetUserID(c)

	//-----------------------------------
	// Get Plan
	//-----------------------------------

	plan, err := h.PlanRepo.GetByID(
		req.PlanID,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "plan not found"},
		)
		return
	}

	//-----------------------------------
	// Validate Amount
	//-----------------------------------

	if req.Amount < plan.MinAmount ||
		req.Amount > plan.MaxAmount {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "amount outside plan limits",
			},
		)
		return
	}

	//-----------------------------------
	// Check Wallet
	//-----------------------------------

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
			gin.H{
				"error": "insufficient balance",
			},
		)
		return
	}

	//-----------------------------------
	// Deduct Balance
	//-----------------------------------

	err = h.WalletRepo.DeductBalance(
		userID,
		req.Amount,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	//-----------------------------------
	// Create Investment
	//-----------------------------------

	startDate := time.Now()

	endDate := startDate.AddDate(
		0,
		0,
		plan.DurationDay,
	)

	investment := &models.Investment{
		UserID:    userID,
		PlanID:    plan.ID,
		Amount:    req.Amount,
		DailyRate: plan.DailyRate,
		StartDate: startDate,
		EndDate:   endDate,
		Status:    "active",
	}

	err = h.InvestmentRepo.CreateInvestment(
		investment,
	)

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
			"message": "investment created",
			"plan":    plan.Name,
			"amount":  req.Amount,
		},
	)
}
func (h *InvestmentHandler) History(
	c *gin.Context,
) {

	userID := utils.GetUserID(c)

	investments, err :=
		h.InvestmentRepo.GetByUserID(
			userID,
		)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		investments,
	)
}
