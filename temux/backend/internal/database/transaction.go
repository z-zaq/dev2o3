package database

import "database/sql"

func WithTransaction(
	db *sql.DB,
	fn func(*sql.Tx) error,
) error {

	tx, err := db.Begin()

	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
