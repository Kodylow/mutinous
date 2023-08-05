package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	// TODO: remove hardcode
	var err error
	db, err = sql.Open("sqlite3", "db/lnaddrServer.db")
	if err != nil {
		return nil, err
	}

	// Create table if not exists
	createTableStatement := `
	CREATE TABLE IF NOT EXISTS users (
		username TEXT PRIMARY KEY,
		balance INTEGER DEFAULT 0
	);`

	_, err = db.Exec(createTableStatement)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetDB() *sql.DB {
	return db
}
