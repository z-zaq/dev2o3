package models

import "time"

type Payment struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Reference  string    `json:"reference"`
	Amount     float64   `json:"amount"`
	Gateway    string    `json:"gateway"`
	Status     string    `json:"status"`
	Verified   bool      `json:"verified"`
	CreatedAt  time.Time `json:"created_at"`
	VerifiedAt time.Time `json:"verified_at"`
}
