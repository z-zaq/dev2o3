package repository

import (
	"database/sql"
	"temux/internal/models"
)

type ReferralRepository struct {
	DB *sql.DB
}

func (r *ReferralRepository) CreateReferral(
	referrerID int,
	referredID int,
) error {

	query := `
	INSERT INTO referrals(
		referrer_id,
		referred_id
	)
	VALUES(?,?)
	`

	_, err := r.DB.Exec(
		query,
		referrerID,
		referredID,
	)

	return err
}
func (r *ReferralRepository) GetByReferrerID(
	referrerID int,
) ([]models.Referral, error) {

	rows, err := r.DB.Query(`
	SELECT
	id,
	referrer_id,
	referred_id,
	created_at
	FROM referrals
	WHERE referrer_id = ?
	ORDER BY created_at DESC
	`, referrerID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var referrals []models.Referral

	for rows.Next() {

		var referral models.Referral

		err := rows.Scan(
			&referral.ID,
			&referral.ReferrerID,
			&referral.ReferredID,
			&referral.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		referrals = append(
			referrals,
			referral,
		)
	}

	return referrals, nil
}
func (r *ReferralRepository) CountReferrals(
	referrerID int,
) (int, error) {

	var count int

	query := `
	SELECT COUNT(*)
	FROM referrals
	WHERE referrer_id = ?
	`

	err := r.DB.QueryRow(
		query,
		referrerID,
	).Scan(&count)

	return count, err
}
