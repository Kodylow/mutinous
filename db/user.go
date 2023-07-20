package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Username string `db:"username"`
	Balance  int    `db:"balance"`
}

func StoreUser(username string) error {
	db := GetDB()
	log.Printf("Storing user with username: %s", username)
	// Check if the user already exists
	if !UserIsInDB(username) {
		// If the user doesn't exist, create a new user
		_, err := db.Exec("INSERT INTO users (username, balance) VALUES (?, ?)", username, 0)
		if err != nil {
			log.Printf("Error when trying to store user: %v", err)
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
	err := db.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(&existingUser.Username, &existingUser.Balance)

	if err != nil {
		log.Println("Error when trying to check if user exists: ", err)
		if err != sql.ErrNoRows {
			return false
		}
		return false
	}

	return true
}
