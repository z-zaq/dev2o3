package handlers

import (
	"net/http"

	"temux/internal/repository"

	"github.com/gin-gonic/gin"
)

type PlanHandler struct {
	Repo *repository.PlanRepository
}

func (h *PlanHandler) GetPlans(
	c *gin.Context,
) {

	plans, err := h.Repo.GetAll()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		plans,
	)
}
