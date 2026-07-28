package repository

import (
	"database/sql"

	"temux/internal/models"
)

type AuditRepository struct {
	DB *sql.DB
}

func (r *AuditRepository) Create(
	log *models.AuditLog,
) error {

	query := `
	INSERT INTO audit_logs(
		user_id,
		action,
		entity,
		entity_id,
		ip_address,
		user_agent
	)
	VALUES(?,?,?,?,?,?)
	`

	_, err := r.DB.Exec(
		query,
		log.UserID,
		log.Action,
		log.Entity,
		log.EntityID,
		log.IPAddress,
		log.UserAgent,
	)

	return err
}

func (r *AuditRepository) GetByUser(
	userID int,
) ([]models.AuditLog, error) {

	rows, err := r.DB.Query(`
	SELECT
		id,
		user_id,
		action,
		entity,
		entity_id,
		ip_address,
		user_agent,
		created_at
	FROM audit_logs
	WHERE user_id = ?
	ORDER BY created_at DESC
	`, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var logs []models.AuditLog

	for rows.Next() {

		var log models.AuditLog

		err := rows.Scan(
			&log.ID,
			&log.UserID,
			&log.Action,
			&log.Entity,
			&log.EntityID,
			&log.IPAddress,
			&log.UserAgent,
			&log.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		logs = append(logs, log)
	}

	return logs, nil
}
