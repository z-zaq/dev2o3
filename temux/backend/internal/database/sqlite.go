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
	CREATE TABLE IF NOT EXISTS wallets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER UNIQUE NOT NULL,
    balance REAL DEFAULT 0,

    FOREIGN KEY(user_id)
    REFERENCES users(id)
);`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return db, nil
}