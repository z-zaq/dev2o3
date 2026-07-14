package repository

import (
	"database/sql"
	"temux/internal/models"
)

type WalletRepository struct {
	DB *sql.DB
}
func (r *WalletRepository) CreateWallet(userID int) error {

	query := `
	INSERT INTO wallets(user_id,balance)
	VALUES(?,0)
	`

	_, err := r.DB.Exec(query, userID)

	return err
}
func (r *WalletRepository) GetWalletByUserID(
	userID int,
) (*models.Wallet, error) {

	wallet := &models.Wallet{}

	query := `
	SELECT id,user_id,balance
	FROM wallets
	WHERE user_id=?
	`

	err := r.DB.QueryRow(
		query,
		userID,
	).Scan(
		&wallet.ID,
		&wallet.UserID,
		&wallet.Balance,
	)

	if err != nil {
		return nil, err
	}

	return wallet, nil
}
func (r *WalletRepository) UpdateBalance(
	userID int,
	amount float64,
) error {

	query := `
	UPDATE wallets
	SET balance=?
	WHERE user_id=?
	`

	_, err := r.DB.Exec(
		query,
		amount,
		userID,
	)

	return err
}