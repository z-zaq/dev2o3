package handlers

import (
	"net/http"

	"temux/internal/repository"
	"temux/internal/services"
	"temux/internal/utils"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	UserRepo *repository.UserRepository
}

func (h *PaymentHandler) InitializeDeposit(
	c *gin.Context,
) {

	var req struct {
		Amount float64 `json:"amount"`
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

	if req.Amount <= 0 {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid amount",
			},
		)

		return
	}

	userID := utils.GetUserID(c)

	user, err := h.UserRepo.GetByID(
		userID,
	)

	if err != nil {

		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "user not found",
			},
		)

		return
	}

	response, err := services.InitializePayment(
		user.Email,
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

	c.JSON(
		http.StatusOK,
		gin.H{
			"authorization_url": response.Data.AuthorizationURL,
			"reference":          response.Data.Reference,
			"access_code":        response.Data.AccessCode,
		},
	)
}
