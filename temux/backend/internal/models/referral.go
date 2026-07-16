package models

import "time"

type Referral struct {
	ID         int       `json:"id"`
	ReferrerID int       `json:"referrer_id"`
	ReferredID int       `json:"referred_id"`
	CreatedAt  time.Time `json:"created_at"`
}
