package services

import (
	"database/sql"
	"fmt"

	"temux/internal/repository"
)

type WalletService struct {
	DB                 *sql.DB
	WalletRepo         *repository.WalletRepository
	TransactionRepo    *repository.TransactionRepository
}

func (s *WalletService) Deposit(
	userID int,
	amount float64,
) (err error) {

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	err = s.WalletRepo.AddBalanceTx(
		tx,
		userID,
		amount,
	)
	if err != nil {
		return err
	}

	err = s.TransactionRepo.CreateDeposit(
		tx,
		userID,
		amount,
	)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
func (s *WalletService) Withdraw(
	userID int,
	amount float64,
) (err error) {

	wallet, err := s.WalletRepo.GetWalletByUserID(
		userID,
	)
	if err != nil {
		return err
	}

	if wallet.Balance < amount {
		return fmt.Errorf("insufficient balance")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	err = s.WalletRepo.DeductBalanceTx(
		tx,
		userID,
		amount,
	)
	if err != nil {
		return err
	}

	err = s.TransactionRepo.CreateWithdrawal(
		tx,
		userID,
		amount,
	)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}