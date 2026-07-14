package handlers

import (
	"net/http"

	"temux/internal/repository"
	"temux/internal/utils"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	WalletRepo *repository.WalletRepository
}
func (h *WalletHandler) GetWallet(
	c *gin.Context,
) {

	userID := utils.GetUserID(c)

	wallet, err := h.WalletRepo.
		GetWalletByUserID(userID)

	if err != nil {

		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "wallet not found",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		wallet,
	)
}