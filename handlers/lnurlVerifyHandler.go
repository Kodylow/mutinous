package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kodylow/mutinous/lightning"
)

// InvoiceStatus represents the status of a payment invoice
type InvoiceStatus struct {
	Status   string `json:"status"`
	Settled  bool   `json:"settled"`
	Preimage string `json:"preimage,omitempty"`
	Pr       string `json:"pr,omitempty"`
}

// LNURLVerifyHandler handles LNURL verification requests
func LNURLVerifyHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	label := params["label"]
	
	// check if invoice has been paid
	invoice, err := lightning.GetInvoiceByLabel(label)
	if err != nil {
		http.Error(w, "Error getting invoice", http.StatusInternalServerError)
		return
	}

	if invoice.Status != "paid" {
		response := InvoiceStatus{
			Status:  "OK",
			Settled: false,
			Preimage: "",
			Pr:      invoice.Bolt11,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := InvoiceStatus{
		Status:   "OK",
		Settled:  true,
		Preimage: invoice.PaymentPreImage, // adjust this as per the actual field name for Preimage in the Invoice struct
		Pr:       invoice.Bolt11,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
