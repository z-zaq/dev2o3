package repository

import (
	"database/sql"
	"temux/internal/models"
)

type WithdrawalRepository struct {
	DB *sql.DB
}

func (r *WithdrawalRepository) Create(
	withdrawal *models.Withdrawal,
) error {

	query := `
	INSERT INTO withdrawals(
		user_id,
		amount,
		status
	)
	VALUES(?,?,?)
	`

	_, err := r.DB.Exec(
		query,
		withdrawal.UserID,
		withdrawal.Amount,
		withdrawal.Status,
	)

	return err
}
func (r *WithdrawalRepository) GetPending() (
	[]models.Withdrawal,
	error,
) {

	rows, err := r.DB.Query(`
	SELECT
	id,
	user_id,
	amount,
	status,
	created_at
	FROM withdrawals
	WHERE status='pending'
	ORDER BY created_at ASC
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var withdrawals []models.Withdrawal

	for rows.Next() {

		var w models.Withdrawal

		err := rows.Scan(
			&w.ID,
			&w.UserID,
			&w.Amount,
			&w.Status,
			&w.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		withdrawals = append(
			withdrawals,
			w)
	}

	return withdrawals, nil
}
func (r *WithdrawalRepository) UpdateStatus(
	id int,
	status string,
) error {

	query := `
	UPDATE withdrawals
	SET status=?
	WHERE id=?
	`

	_, err := r.DB.Exec(
		query,
		status,
		id,
	)

	return err
}
func (r *WithdrawalRepository) GetByID(
	id int,
) (*models.Withdrawal, error) {

	withdrawal := &models.Withdrawal{}

	query := `
	SELECT
		id,
		user_id,
		amount,
		status,
		created_at
	FROM withdrawals
	WHERE id = ?
	`

	err := r.DB.QueryRow(
		query,
		id,
	).Scan(
		&withdrawal.ID,
		&withdrawal.UserID,
		&withdrawal.Amount,
		&withdrawal.Status,
		&withdrawal.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return withdrawal, nil
}
