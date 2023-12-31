package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/elementsproject/glightning/glightning"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kodylow/mutinous/db"
	"github.com/kodylow/mutinous/lightning"
	"github.com/kodylow/mutinous/utils"
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
	username = strings.TrimSpace(username)
	log.Println("LNURL callback from ", username)

	// check if user exists
	if !db.UserIsInDB(username) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ERROR", "reason": "User is not yet Mutinous"})
		return
	}

	log.Println("User exists...")

	// check if amount is valid
	amountStr := r.URL.Query().Get("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil || amount < 1000 || amount > 110000 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ERROR", "reason": "Invalid amount"})
		return
	}

	// create a label just username and a random base64 string
	label, err := utils.GenerateLabel(username)
	log.Println("Getting invoice for label: ", label)
	invoice, err := lightning.CreateInvoice(uint64(amount), label, utils.GetMetadata(username))

	log.Println("Invoice created...")

	resp := buildCallbackResponse(username, invoice, label)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// buildCallbackResponse builds the response to a LNURL callback
func buildCallbackResponse(username string, invoice *glightning.Invoice, label string) *LNURLCallbackResponse {
	resp := &LNURLCallbackResponse{
		Status: "OK",
	}
	resp.SuccessAction.Tag = "message"
	resp.SuccessAction.Message = "Walk the plank, this Mutiny's just getting started!"
	resp.Verify = DOMAIN + "/lnurlp/" + username + "/verify/" + label
	resp.Pr = invoice.Bolt11
	return resp
}
