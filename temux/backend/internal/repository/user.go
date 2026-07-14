package repository

import (
	"database/sql"
	"temux/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user *models.User) error {

	query := `
	INSERT INTO users
	(name,email,password,referral_code)
	VALUES(?,?,?,?)
	`

	_, err := r.DB.Exec(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.ReferralCode,
	)

	return err
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {

	user := &models.User{}

	query := `
	SELECT id,name,email,password,
	referral_code,is_admin
	FROM users
	WHERE email=?
	`

	err := r.DB.QueryRow(query, email).Scan(
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