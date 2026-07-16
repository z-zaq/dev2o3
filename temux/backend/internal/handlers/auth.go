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
	Repo         *repository.UserRepository
	WalletRepo   *repository.WalletRepository
	ReferralRepo *repository.ReferralRepository
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
		Name         string `json:"name"`
		Email        string `json:"email"`
		Password     string `json:"password"`
		ReferralCode string `json:"referral_code"`
	}

	//-----------------------------------
	// Parse Request
	//-----------------------------------

	if err := c.BindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	//-----------------------------------
	// Hash Password
	//-----------------------------------

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to hash password"},
		)
		return
	}

	//-----------------------------------
	// Create User Object
	//-----------------------------------

	user := &models.User{
		Name:         req.Name,
		Email:        req.Email,
		Password:     string(hash),
		ReferralCode: generateReferral(),
	}

	//-----------------------------------
	// Create User
	//-----------------------------------

	userID, err := h.Repo.CreateUser(user)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	//-----------------------------------
	// Create Wallet
	//-----------------------------------

	err = h.WalletRepo.CreateWallet(
		int(userID),
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "wallet creation failed"},
		)
		return
	}

	//-----------------------------------
	// Process Referral
	//-----------------------------------

	if req.ReferralCode != "" {

		referrer, err := h.Repo.GetByReferralCode(
			req.ReferralCode,
		)

		if err == nil {

			err = h.ReferralRepo.CreateReferral(
				referrer.ID,
				int(userID),
			)

			if err != nil {
				c.JSON(
					http.StatusInternalServerError,
					gin.H{
						"error": "referral creation failed",
					},
				)
				return
			}
		}
	}

	//-----------------------------------
	// Success
	//-----------------------------------

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message":       "registration successful",
			"user_id":       userID,
			"referral_code": user.ReferralCode,
		},
	)
}
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
