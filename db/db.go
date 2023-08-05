package db

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type User struct {
	Username string `db:"username"`
	Balance  int    `db:"balance"`
}

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

func StoreUser(username string) error {
	if db == nil {
		return errors.New("Database is not initialized")
	}

	log.Printf("Storing user with username: %s", username)
	if !UserIsInDB(username) {
		_, err := db.Exec("INSERT INTO users (username, balance) VALUES (?, ?)", username, 0)
		if err != nil {
			log.Printf("Error when trying to store user: %v", err)
			return err
		}
	}
	return nil
}

func UserIsInDB(username string) bool {
	if db == nil {
		log.Fatal("Database is not initialized")
		return false
	}

	var existingUser User
	err := db.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(&existingUser.Username, &existingUser.Balance)

	if err != nil {
		if err == sql.ErrNoRows {
			return false // <-- User not found, return false
		}
		log.Println("Error when trying to check if user exists:", err)
		return false
	}

	return true
}


func GetUserBalance(username string) (int, error) {
	if db == nil {
		return 0, errors.New("Database is not initialized")
	}

	var balance int
	err := db.QueryRow("SELECT balance FROM users WHERE username = ?", username).Scan(&balance)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("user not found")
		}
		log.Printf("Error when trying to retrieve user balance: %v", err)
		return 0, err
	}

	return balance, nil
}

func AddToUserBalance(username string, msat int) error {
	if db == nil {
		return errors.New("Database is not initialized")
	}

	_, err := db.Exec("UPDATE users SET balance = balance + ? WHERE username = ?", msat, username)
	if err != nil {
		log.Printf("Error when trying to add to user's balance: %v", err)
		return err
	}
	return nil
}

func DeductFromUserBalance(username string, amount int) error {
	if db == nil {
		return errors.New("Database is not initialized")
	}

	currentBalance, err := GetUserBalance(username)
	if err != nil {
		return err
	}

	if currentBalance < amount {
		return errors.New("insufficient funds")
	}

	_, err = db.Exec("UPDATE users SET balance = balance - ? WHERE username = ?", amount, username)
	if err != nil {
		log.Printf("Error when trying to deduct from user's balance: %v", err)
		return err
	}
	return nil
}
