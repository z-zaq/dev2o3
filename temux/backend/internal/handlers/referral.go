package handlers

import (
	"net/http"
	"temux/internal/repository"
	"temux/internal/utils"

	"github.com/gin-gonic/gin"
)

type ReferralHandler struct {
	ReferralRepo *repository.ReferralRepository
}

func (h *ReferralHandler) MyReferrals(
	c *gin.Context,
) {

	userID := utils.GetUserID(c)

	referrals, err :=
		h.ReferralRepo.GetByReferrerID(
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
		referrals,
	)
}
func (h *ReferralHandler) Stats(
	c *gin.Context,
) {

	userID := utils.GetUserID(c)

	total, err :=
		h.ReferralRepo.CountReferrals(
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
		gin.H{
			"total_referrals": total,
		},
	)
}
