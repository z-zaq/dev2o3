package models

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"-"`
	ReferralCode string `json:"referral_code"`
	IsAdmin      bool   `json:"is_admin"`
}