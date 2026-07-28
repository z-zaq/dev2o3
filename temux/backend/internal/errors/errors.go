package errors

import "errors"

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrInvalidAmount       = errors.New("invalid amount")
	ErrPlanNotFound        = errors.New("plan not found")
	ErrPaymentNotVerified  = errors.New("payment not verified")
	ErrInvestmentNotFound  = errors.New("investment not found")
	ErrWithdrawalPending   = errors.New("withdrawal already pending")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")
)
