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

	//-----------------------------------
	// Users Table
	//-----------------------------------

	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		referral_code TEXT UNIQUE,
		is_admin BOOLEAN DEFAULT FALSE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = db.Exec(usersTable)
	if err != nil {
		return nil, err
	}

	//-----------------------------------
	// Wallets Table
	//-----------------------------------

	walletsTable := `
	CREATE TABLE IF NOT EXISTS wallets (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER UNIQUE NOT NULL,
		balance REAL DEFAULT 0,

		FOREIGN KEY(user_id)
		REFERENCES users(id)
	);
	`

	_, err = db.Exec(walletsTable)
	if err != nil {
		return nil, err
	}

	return db, nil
}
