package handlers

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/kodylow/mutinous/db"
)

type Payment struct {
	Label         string `json:"label"`
	PreImage      string `json:"preimage"`
	MilliSatoshis string `json:"msat"`
}

func InvoicePaidHandler(w http.ResponseWriter, r *http.Request) {
	host, port, err := net.SplitHostPort(r.RemoteAddr)
	log.Println("Host:", host, "\nPort:", port)
	if err != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if host != "127.0.0.1" && host != "::1" && host != "localhost" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var paymentDetails Payment // Define a struct for the expected payment details

	err = json.NewDecoder(r.Body).Decode(&paymentDetails)
	if err != nil {
		log.Printf("Error decoding JSON: %s", err) // <--- Add detailed log here
		http.Error(w, "Bad Request 1", http.StatusBadRequest)
		return
	}

	// Parse off username from label "username-invoiceID"
	username := strings.TrimSpace(strings.Split(paymentDetails.Label, "-")[0])
	log.Println("Username:", username)
	// Get the user from the database
	yes := db.UserIsInDB(username)
	if !yes {
		// User not found, add them to the database with a balance of 0
		err := db.StoreUser(username)
		if err != nil {
			log.Printf("Error storing new user: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	// Add the payment millisatoshi amount to the user's balance, need to convert to int
	msat, err := strconv.Atoi(paymentDetails.MilliSatoshis)
	if err != nil {
		log.Printf("Error converting MilliSatoshis to int: %s", err) // <--- Add detailed log here
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	userBalance, err := db.GetUserBalance(username)
	if err != nil {
		log.Println("Error finding user in DB: ", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	log.Println("User balance: ", userBalance)

	db.AddToUserBalance(username, msat)

	userBalance, err = db.GetUserBalance(username)
	if err != nil {
		log.Printf("Error retrieving user balance: %s", err) // <--- Add detailed log here
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	log.Println("User balance after: ", userBalance)

	w.Write([]byte("OK"))
}
