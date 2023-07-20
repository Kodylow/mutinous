package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name    string `db:"name"`
	Balance int    `db:"balance"`
}

func StoreUser(username string) error {
	db := GetDB()

	// Check if the user already exists
	if !UserIsInDB(username) {
		// If the user doesn't exist, create a new user
		_, err := db.Exec("INSERT INTO users (name, balance) VALUES (?, ?)", username, 0)
		if err != nil {
			return err
		}
		return nil
	}

	// If the user exists, return nil
	return nil
}

func UserIsInDB(username string) bool {
	db := GetDB()

	var existingUser User
	err := db.QueryRow("SELECT * FROM users WHERE name = ?", username).Scan(&existingUser)
	if err != nil {
		if err != sql.ErrNoRows {
			return false
		}
		return false
	}

	return true
}