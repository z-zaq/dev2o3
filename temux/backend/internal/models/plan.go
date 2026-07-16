package models

type Plan struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	MinAmount   float64 `json:"min_amount"`
	MaxAmount   float64 `json:"max_amount"`
	DailyRate   float64 `json:"daily_rate"`
	DurationDay int     `json:"duration_day"`
}
