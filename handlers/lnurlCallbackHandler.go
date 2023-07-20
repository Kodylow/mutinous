package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kodylow/mutinous/db"
	"github.com/kodylow/mutinous/lightning"
)

// LNURLCallbackResponse is the response to a LNURL callback
type LNURLCallbackResponse struct {
	Status        string `json:"status"`
	SuccessAction struct {
		Tag     string `json:"tag"`
		Message string `json:"message"`
	} `json:"successAction"`
	Verify string   `json:"verify"`
	Routes []string `json:"routes"`
	Pr     string   `json:"pr"`
}

// LNURLCallbackHandler handles LNURL callbacks
func LNURLCallbackHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	
	// check if user exists
	userInDB := db.UserIsInDB(username)
	if !userInDB {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ERROR", "reason": "User is not yet Mutinous"})
		return
	}

	// check if amount is valid
	amountStr := r.URL.Query().Get("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil || amount < 1000 || amount > 110000 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ERROR", "reason": "Invalid amount"})
		return
	}

	// grab invoice for the amount
	invoice, err := lightning.CreateInvoice(uint64(amount), username, "Sats for "+username, 3600, nil, "", false)
	resp := &LNURLCallbackResponse{
		Status: "OK",
	}
	resp.SuccessAction.Tag = "message"
	resp.SuccessAction.Message = "Walk the plank, this Mutiny's just getting started!"
	resp.Verify = "https://getalby.com/lnurlp/" + username + "/verify/ch4Z7u3xYo5tWWSGsafLVHqZ"
	resp.Pr = invoice.Bolt11

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}