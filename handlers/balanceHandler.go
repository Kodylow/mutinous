package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kodylow/mutinous/db"
)

type BalanceResponse struct {
	Username string `json:"username"`
	Balance  int    `json:"balance"`
	Error    string `json:"error,omitempty"`
}

func BalanceHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is from localhost
	if r.RemoteAddr != "127.0.0.1:PORT" && r.RemoteAddr != "[::1]:PORT" {
		http.Error(w, "Forbidden: Access allowed only from localhost", http.StatusForbidden)
		return
	}

	params := mux.Vars(r)
	username := params["username"]

	// Fetch balance from the database
	balance, err := db.GetUserBalance(username)
	if err != nil {
		http.Error(w, "Error fetching balance", http.StatusInternalServerError)
		return
	}

	resp := &BalanceResponse{
		Username: username,
		Balance:  balance,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
