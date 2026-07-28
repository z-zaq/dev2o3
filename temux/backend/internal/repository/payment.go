package repository

import (
	"database/sql"
	"time"

	"temux/internal/models"
)

type PaymentRepository struct {
	DB *sql.DB
}

func (r *PaymentRepository) Create(
	payment *models.Payment,
) error {

	query := `
	INSERT INTO payments(
		user_id,
		reference,
		amount,
		gateway,
		status,
		verified
	)
	VALUES(?,?,?,?,?,?)
	`

	_, err := r.DB.Exec(
		query,
		payment.UserID,
		payment.Reference,
		payment.Amount,
		payment.Gateway,
		payment.Status,
		payment.Verified,
	)

	return err
}

func (r *PaymentRepository) GetByReference(
	reference string,
) (*models.Payment, error) {

	payment := &models.Payment{}

	query := `
	SELECT
		id,
		user_id,
		reference,
		amount,
		gateway,
		status,
		verified,
		created_at,
		verified_at
	FROM payments
	WHERE reference = ?
	`

	err := r.DB.QueryRow(
		query,
		reference,
	).Scan(
		&payment.ID,
		&payment.UserID,
		&payment.Reference,
		&payment.Amount,
		&payment.Gateway,
		&payment.Status,
		&payment.Verified,
		&payment.CreatedAt,
		&payment.VerifiedAt,
	)

	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (r *PaymentRepository) MarkVerified(
	id int,
) error {

	query := `
	UPDATE payments
	SET
		verified = 1,
		status = 'completed',
		verified_at = ?
	WHERE id = ?
	`

	_, err := r.DB.Exec(
		query,
		time.Now(),
		id,
	)

	return err
}

func (r *PaymentRepository) Exists(
	reference string,
) (bool, error) {

	var count int

	query := `
	SELECT COUNT(*)
	FROM payments
	WHERE reference = ?
	`

	err := r.DB.QueryRow(
		query,
		reference,
	).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
