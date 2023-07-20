package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./lightning.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetDB() *sql.DB {
	return db
}
