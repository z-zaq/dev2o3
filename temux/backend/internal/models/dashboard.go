package models

type Dashboard struct {
	WalletBalance     float64 `json:"wallet_balance"`
	TotalDeposits     float64 `json:"total_deposits"`
	TotalWithdrawals  float64 `json:"total_withdrawals"`
	ActiveInvestments int     `json:"active_investments"`
	TotalInvested     float64 `json:"total_invested"`
	TotalProfit       float64 `json:"total_profit"`
}
