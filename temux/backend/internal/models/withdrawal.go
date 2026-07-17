package models

import "time"

type Withdrawal struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
