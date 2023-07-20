package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kodylow/mutinous/db"
	"github.com/kodylow/mutinous/utils"
)

type LNAddressResponse struct {
	Status         string `json:"status"`
	Tag            string `json:"tag"`
	CommentAllowed int    `json:"commentAllowed"`
	Callback       string `json:"callback"`
	Metadata       string `json:"metadata"`
	MinSendable    int    `json:"minSendable"`
	MaxSendable    int    `json:"maxSendable"`
	PayerData      struct {
		Name struct {
			Mandatory bool `json:"mandatory"`
		} `json:"name"`
		Email struct {
			Mandatory bool `json:"mandatory"`
		} `json:"email"`
		Pubkey struct {
			Mandatory bool `json:"mandatory"`
		} `json:"pubkey"`
	} `json:"payerData"`
	NostrPubkey string `json:"nostrPubkey"`
	AllowsNostr bool   `json:"allowsNostr"`
}

func LNAddressHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]

	// check if user exists on replit
	if !utils.ValidateReplitUser(username) {
		http.Error(w, "Invalid username", http.StatusBadRequest)
		return
	}

	// store user in db if not exists
	err := db.StoreUser(username)
	if err != nil {
		http.Error(w, "Unable to create user", http.StatusInternalServerError)
		return
	}

	resp := &LNAddressResponse{
		Status:         "OK",
		Tag:            "payRequest",
		CommentAllowed: 255,
		Callback:       "https://domain.com/lnurlp/" + username + "/callback",
		Metadata:       "[[\"text/identifier\",\"" + username + "@domain.com\"],[\"text/plain\",\"Sats for " + username + "\"]]",
		MinSendable:    1000,
		MaxSendable:    110000,
	}
	resp.PayerData.Name.Mandatory = false
	resp.PayerData.Email.Mandatory = false
	resp.PayerData.Pubkey.Mandatory = false
	resp.NostrPubkey = ""
	resp.AllowsNostr = false

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}