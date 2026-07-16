package repository

import (
	"database/sql"

	"temux/internal/models"
)

type TransactionRepository struct {
	DB *sql.DB
}

func (r *TransactionRepository) CreateTransaction(
	tx *models.Transaction,
) error {

	query := `
	INSERT INTO transactions(
		user_id,
		type,
		amount
	)
	VALUES(?,?,?)
	`

	_, err := r.DB.Exec(
		query,
		tx.UserID,
		tx.Type,
		tx.Amount,
	)

	return err
}
func (r *TransactionRepository) GetByUserID(
	userID int,
) ([]models.Transaction, error) {

	query := `
	SELECT
		id,
		user_id,
		type,
		amount,
		created_at
	FROM transactions
	WHERE user_id = ?
	ORDER BY created_at DESC
	`

	rows, err := r.DB.Query(
		query,
		userID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var transactions []models.Transaction

	for rows.Next() {

		var tx models.Transaction

		err := rows.Scan(
			&tx.ID,
			&tx.UserID,
			&tx.Type,
			&tx.Amount,
			&tx.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		transactions = append(
			transactions,
			tx,
		)
	}

	return transactions, nil
}
func (r *TransactionRepository) GetTotalDeposits(
	userID int,
) (float64, error) {

	var total float64

	query := `
	SELECT COALESCE(SUM(amount),0)
	FROM transactions
	WHERE user_id = ?
	AND type = 'deposit'
	`

	err := r.DB.QueryRow(
		query,
		userID,
	).Scan(&total)

	return total, err
}
func (r *TransactionRepository) GetTotalWithdrawals(
	userID int,
) (float64, error) {

	var total float64

	query := `
	SELECT COALESCE(SUM(amount),0)
	FROM transactions
	WHERE user_id = ?
	AND type = 'withdraw'
	`

	err := r.DB.QueryRow(
		query,
		userID,
	).Scan(&total)

	return total, err
}
