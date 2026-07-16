package handlers

import (
	"math/rand"
	"net/http"

	"temux/internal/auth"
	"temux/internal/models"
	"temux/internal/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	Repo       *repository.UserRepository
	WalletRepo *repository.WalletRepository
}

func generateReferral() string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	code := "TEMUX-"

	for i := 0; i < 6; i++ {
		code += string(
			chars[rand.Intn(len(chars))],
		)
	}

	return code
}

func (h *AuthHandler) Register(c *gin.Context) {

	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	user := &models.User{
		Name:         req.Name,
		Email:        req.Email,
		Password:     string(hash),
		ReferralCode: generateReferral(),
	}

	userID, err := h.Repo.CreateUser(user)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	err = h.WalletRepo.CreateWallet(
		int(userID),
	)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "failed to create wallet",
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "registered",
		},
	)
}

// 	user := &models.User{
// 		Name:         req.Name,
// 		Email:        req.Email,
// 		Password:     string(hash),
// 		ReferralCode: generateReferral(),
// 	}

// 	userID, err := h.Repo.CreateUser(user)
// 	if err != nil {
// 	c.JSON(
// 		http.StatusBadRequest,
// 		gin.H{"error": err.Error()},
// 	)
// 	return
// }

// 	c.JSON(http.StatusCreated,
// 		gin.H{"message": "registered"})
// }

func (h *AuthHandler) Login(c *gin.Context) {

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	c.BindJSON(&req)

	user, err := h.Repo.GetByEmail(req.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized,
			gin.H{"error": "invalid credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized,
			gin.H{"error": "invalid credentials"})
		return
	}

	token, _ := auth.GenerateToken(user.ID)

	c.JSON(http.StatusOK,
		gin.H{
			"token": token,
			"user":  user,
		})
}
