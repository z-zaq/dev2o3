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
func (r *ReferralRepository) GetReferrerByReferredID(
	referredID int,
) (int, error) {

	var referrerID int

	query := `
	SELECT referrer_id
	FROM referrals
	WHERE referred_id = ?
	`

	err := r.DB.QueryRow(
		query,
		referredID,
	).Scan(&referrerID)

	return referrerID, err
}
func (r *ReferralRepository) CreateReward(
	reward *models.ReferralReward,
) error {

	query := `
	INSERT INTO referral_rewards(
		referrer_id,
		referred_id,
		deposit_amount,
		commission
	)
	VALUES(?,?,?,?)
	`

	_, err := r.DB.Exec(
		query,
		reward.ReferrerID,
		reward.ReferredID,
		reward.DepositAmount,
		reward.Commission,
	)

	return err
}
func (r *ReferralRepository) GetRewardsByReferrerID(
	referrerID int,
) ([]models.ReferralReward, error) {

	rows, err := r.DB.Query(`
	SELECT
	id,
	referrer_id,
	referred_id,
	deposit_amount,
	commission,
	created_at
	FROM referral_rewards
	WHERE referrer_id = ?
	ORDER BY created_at DESC
	`, referrerID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rewards []models.ReferralReward

	for rows.Next() {

		var reward models.ReferralReward

		err := rows.Scan(
			&reward.ID,
			&reward.ReferrerID,
			&reward.ReferredID,
			&reward.DepositAmount,
			&reward.Commission,
			&reward.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		rewards = append(
			rewards,
			reward,
		)
	}

	return rewards, nil
}
func (r *ReferralRepository) TotalReferralEarnings(
	referrerID int,
) (float64, error) {

	var total float64

	query := `
	SELECT COALESCE(
		SUM(commission),
		0
	)
	FROM referral_rewards
	WHERE referrer_id = ?
	`

	err := r.DB.QueryRow(
		query,
		referrerID,
	).Scan(&total)

	return total, err
}
func (r *ReferralRepository) TotalReferrals() (
	int,
	error,
) {

	var count int

	query := `
	SELECT COUNT(*)
	FROM referrals
	`

	err := r.DB.QueryRow(
		query,
	).Scan(&count)

	return count, err
}
