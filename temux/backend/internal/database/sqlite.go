package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "temux.db")
	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		referral_code TEXT,
		is_admin BOOLEAN DEFAULT FALSE
	);
	`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return db, nil
}