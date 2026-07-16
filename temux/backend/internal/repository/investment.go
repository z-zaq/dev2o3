package repository

import (
	"database/sql"

	"temux/internal/models"
)

type InvestmentRepository struct {
	DB *sql.DB
}

func (r *InvestmentRepository) CreateInvestment(
	investment *models.Investment,
) error {

	query := `
	INSERT INTO investments(
		user_id,
		plan_id,
		amount,
		daily_rate,
		start_date,
		end_date,
		status
	)
	VALUES(?,?,?,?,?,?,?)
	`

	_, err := r.DB.Exec(
		query,
		investment.UserID,
		investment.PlanID,
		investment.Amount,
		investment.DailyRate,
		investment.StartDate,
		investment.EndDate,
		investment.Status,
	)

	return err
}
func (r *InvestmentRepository) GetByUserID(
	userID int,
) ([]models.Investment, error) {

	rows, err := r.DB.Query(`
	SELECT
	id,
	user_id,
	plan_id,
	amount,
	daily_rate,
	start_date,
	end_date,
	status
	FROM investments
	WHERE user_id = ?
	ORDER BY start_date DESC
	`, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var investments []models.Investment

	for rows.Next() {

		var inv models.Investment

		err := rows.Scan(
			&inv.ID,
			&inv.UserID,
			&inv.PlanID,
			&inv.Amount,
			&inv.DailyRate,
			&inv.StartDate,
			&inv.EndDate,
			&inv.Status,
		)

		if err != nil {
			return nil, err
		}

		investments = append(
			investments,
			inv,
		)
	}

	return investments, nil
}
func (r *InvestmentRepository) UpdateProfit(
	id int,
	profit float64,
) error {

	query := `
	UPDATE investments
	SET profit_earned = ?
	WHERE id = ?
	`

	_, err := r.DB.Exec(
		query,
		profit,
		id,
	)

	return err
}
func (r *InvestmentRepository) UpdateClaimedProfit(
	id int,
	amount float64,
) error {

	query := `
	UPDATE investments
	SET claimed_profit = ?
	WHERE id = ?
	`

	_, err := r.DB.Exec(
		query,
		amount,
		id,
	)

	return err
}
func (r *InvestmentRepository) GetByID(
	id int,
) (*models.Investment, error) {

	inv := &models.Investment{}

	query := `
	SELECT
	id,
	user_id,
	plan_id,
	amount,
	daily_rate,
	profit_earned,
	claimed_profit,
	start_date,
	end_date,
	status
	FROM investments
	WHERE id = ?
	`

	err := r.DB.QueryRow(
		query,
		id,
	).Scan(
		&inv.ID,
		&inv.UserID,
		&inv.PlanID,
		&inv.Amount,
		&inv.DailyRate,
		&inv.ProfitEarned,
		&inv.ClaimedProfit,
		&inv.StartDate,
		&inv.EndDate,
		&inv.Status,
	)

	if err != nil {
		return nil, err
	}

	return inv, nil
}
func (r *InvestmentRepository) CountActive(
	userID int,
) (int, error) {

	var count int

	query := `
	SELECT COUNT(*)
	FROM investments
	WHERE user_id = ?
	AND status = 'active'
	`

	err := r.DB.QueryRow(
		query,
		userID,
	).Scan(&count)

	return count, err
}
func (r *InvestmentRepository) TotalInvested(
	userID int,
) (float64, error) {

	var total float64

	query := `
	SELECT COALESCE(SUM(amount),0)
	FROM investments
	WHERE user_id = ?
	`

	err := r.DB.QueryRow(
		query,
		userID,
	).Scan(&total)

	return total, err
}
func (r *InvestmentRepository) TotalProfit(
	userID int,
) (float64, error) {

	var total float64

	query := `
	SELECT COALESCE(SUM(profit_earned),0)
	FROM investments
	WHERE user_id = ?
	`

	err := r.DB.QueryRow(
		query,
		userID,
	).Scan(&total)

	return total, err
}
