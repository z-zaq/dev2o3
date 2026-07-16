package models

import "time"

type ReferralReward struct {
	ID            int       `json:"id"`
	ReferrerID    int       `json:"referrer_id"`
	ReferredID    int       `json:"referred_id"`
	DepositAmount float64   `json:"deposit_amount"`
	Commission    float64   `json:"commission"`
	CreatedAt     time.Time `json:"created_at"`
}
