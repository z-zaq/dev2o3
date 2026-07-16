package models

import "time"

type Investment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PlanID    int       `json:"plan_id"`
	Amount    float64   `json:"amount"`
	DailyRate float64   `json:"daily_rate"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    string    `json:"status"`
}
