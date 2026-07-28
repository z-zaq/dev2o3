package handlers

import (
	"net/http"

	
	"temux/internal/repository"
	"temux/internal/services"

	"github.com/gin-gonic/gin"
)

type PaymentVerificationHandler struct {
	WalletRepo      *repository.WalletRepository
	TransactionRepo *repository.TransactionRepository
}

func (h *PaymentVerificationHandler) VerifyDeposit(
	c *gin.Context,
) {

	var req struct {
		Reference string `json:"reference"`
	}

	if err := c.BindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	verify, err := services.VerifyPayment(
		req.Reference,
	)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	if verify.Data.Status != "success" {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "payment not successful",
			},
		)

		return
	}

	// Wallet crediting logic will be added
	// after we create the payment record table.

	c.JSON(
		http.StatusOK,
		gin.H{
			"message":   "payment verified",
			"reference": verify.Data.Reference,
			"amount":    float64(verify.Data.Amount) / 100,
		},
	)
}
