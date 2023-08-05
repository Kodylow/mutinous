package handlers

import (
	"encoding/json"
	"log"
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
	// Verify the request is from localhost
	if r.RemoteAddr != "127.0.0.1:8080" && r.RemoteAddr != "localhost:8080" { // Replace PORT with the actual port your server listens to
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var paymentDetails Payment // Define a struct for the expected payment details

	err := json.NewDecoder(r.Body).Decode(&paymentDetails)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Parse off username from label "username-invoiceID"
	username := strings.Split(paymentDetails.Label, "-")[0]

	// Get the user from the database
	yes := db.UserIsInDB(username)
	if !yes {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Add the payment millisatoshi amount to the user's balance, need to convert to int
	msat, err := strconv.Atoi(paymentDetails.MilliSatoshis)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	userBalance, err := db.GetUserBalance(username)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	log.Println("User balance: ", userBalance)

	db.AddToUserBalance(username, msat)

	userBalance, err = db.GetUserBalance(username)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	log.Println("User balance after: ", userBalance)

	w.Write([]byte("OK"))
}
