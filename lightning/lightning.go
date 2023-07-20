package lightning

import (
	"github.com/niftynei/glightning/glightning"
)

// Custom wrapper to limit the scope of the lightning client
var lightning *glightning.Lightning

// InitLightning initializes the lightning client
func InitLightning() {
	lightning = glightning.NewLightning()
	lightning.StartUp("lightning-rpc", "/tmp/clight-1")
}

// CreateInvoice creates an invoice
func CreateInvoice(satoshi uint64, label string, description string, expiry uint32) (*glightning.Invoice, error) {
	return lightning.CreateInvoice(satoshi, label, description, expiry, nil, "", false)
}

// IsInvoicePaid checks if an invoice is paid
func IsInvoicePaid(label string) bool {
	invoice, err := lightning.GetInvoice(label)
	if err != nil {
		panic(err)
	}

	// check invoice status
	if invoice.Status != "paid" {
		return false
	}

	return true
}
