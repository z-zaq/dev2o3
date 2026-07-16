package repository

import (
	"database/sql"

	"temux/internal/models"
)

type PlanRepository struct {
	DB *sql.DB
}

func (r *PlanRepository) GetAll() (
	[]models.Plan,
	error,
) {

	rows, err := r.DB.Query(`
	SELECT
	id,
	name,
	min_amount,
	max_amount,
	daily_rate,
	duration_day
	FROM plans
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var plans []models.Plan

	for rows.Next() {

		var p models.Plan

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.MinAmount,
			&p.MaxAmount,
			&p.DailyRate,
			&p.DurationDay,
		)

		if err != nil {
			return nil, err
		}

		plans = append(plans, p)
	}

	return plans, nil
}
func (r *PlanRepository) SeedPlans() error {

	query := `
	INSERT OR IGNORE INTO plans
	(id,name,min_amount,max_amount,daily_rate,duration_day)
	VALUES
	(1,'Starter',100,999,2,30),
	(2,'Premium',1000,4999,3,45),
	(3,'VIP',5000,9999,4,60),
	(4,'Enterprise',10000,100000,5,90)
	`

	_, err := r.DB.Exec(query)

	return err
}
func (r *PlanRepository) GetByID(
	id int,
) (*models.Plan, error) {

	plan := &models.Plan{}

	query := `
	SELECT
		id,
		name,
		min_amount,
		max_amount,
		daily_rate,
		duration_day
	FROM plans
	WHERE id = ?
	`

	err := r.DB.QueryRow(
		query,
		id,
	).Scan(
		&plan.ID,
		&plan.Name,
		&plan.MinAmount,
		&plan.MaxAmount,
		&plan.DailyRate,
		&plan.DurationDay,
	)

	if err != nil {
		return nil, err
	}

	return plan, nil
}
