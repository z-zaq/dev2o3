package repository

import (
	"database/sql"

	"temux/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(
	user *models.User,
) (int64, error) {

	query := `
INSERT INTO users
(name, email, password, referral_code)
VALUES (?, ?, ?, ?)
`

	result, err := r.DB.Exec(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.ReferralCode,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (r *UserRepository) GetByEmail(
	email string,
) (*models.User, error) {

	user := &models.User{}

	query := `
SELECT
	id,
	name,
	email,
	password,
	referral_code,
	is_admin,
	created_at
FROM users
WHERE email = ?
`

	err := r.DB.QueryRow(
		query,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.ReferralCode,
		&user.IsAdmin,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil

}
func (r *UserRepository) GetByReferralCode(
	code string,
) (*models.User, error) {

	user := &models.User{}

	query := `
	SELECT
	id,
	name,
	email,
	password,
	referral_code,
	is_admin
	FROM users
	WHERE referral_code = ?
	`

	err := r.DB.QueryRow(
		query,
		code,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.ReferralCode,
		&user.IsAdmin,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
