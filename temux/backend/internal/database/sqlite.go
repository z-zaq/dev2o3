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

	walletsTable := `
	CREATE TABLE IF NOT EXISTS wallets (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER UNIQUE NOT NULL,
		balance REAL DEFAULT 0,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	_, err = db.Exec(walletsTable)
	if err != nil {
		return nil, err
	}

	transactionsTable := `
	CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		type TEXT NOT NULL,
		amount REAL NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	_, err = db.Exec(transactionsTable)
	if err != nil {
		return nil, err
	}
	plansTable := `
CREATE TABLE IF NOT EXISTS plans (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	min_amount REAL NOT NULL,
	max_amount REAL NOT NULL,
	daily_rate REAL NOT NULL,
	duration_day INTEGER NOT NULL
);
`

	_, err = db.Exec(plansTable)
	if err != nil {
		return nil, err
	}
	investmentsTable := `
CREATE TABLE IF NOT EXISTS investments (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER NOT NULL,
	plan_id INTEGER NOT NULL,
	amount REAL NOT NULL,
	daily_rate REAL NOT NULL,
	profit_earned REAL DEFAULT 0,
	claimed_profit REAL DEFAULT 0,
	start_date DATETIME NOT NULL,
	end_date DATETIME NOT NULL,
	status TEXT DEFAULT 'active',

	FOREIGN KEY(user_id)
	REFERENCES users(id),

	FOREIGN KEY(plan_id)
	REFERENCES plans(id)
);
`

	_, err = db.Exec(investmentsTable)
	if err != nil {
		return nil, err
	}
	referralsTable := `
CREATE TABLE IF NOT EXISTS referrals (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	referrer_id INTEGER NOT NULL,
	referred_id INTEGER NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

	FOREIGN KEY(referrer_id)
	REFERENCES users(id),

	FOREIGN KEY(referred_id)
	REFERENCES users(id)
);
`

	_, err = db.Exec(referralsTable)
	if err != nil {
		return nil, err
	}

	return db, nil
}
