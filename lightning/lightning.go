package lightning

import (
	"bytes"
	"errors"
	"log"
	"os/exec"

	"github.com/elementsproject/glightning/glightning"
)

// Custom wrapper to limit the scope of the lightning client
var lightning *glightning.Lightning

// InitLightning initializes the lightning client
func InitLightning(rpcPath string) error {
	// Test lightning-cli
	cmd := exec.Command("lightning-cli", "--lightning-dir=./.lightning", "--signet", "getinfo")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Println("Failed to run lightning-cli: is lightningd running? Try running `just cln`: ", err)
		return errors.New("lightningd isn't running")
	}

	if out.Len() == 0 {
		log.Println("lightning-cli produced no output.")
		return errors.New("lightningd isn't running")
	}

	log.Println("lightning-cli response:", out.String())

	lightning = glightning.NewLightning()
	log.Println("RPC path:", rpcPath)
	lightning.StartUp("lightning-rpc", rpcPath)

	log.Printf("Lightning client initialized...\n")
	return nil
}

// CreateInvoice creates an invoice
func CreateInvoice(satoshi uint64, label string, description string) (*glightning.Invoice, error) {
	log.Println("Creating invoice...")
	return lightning.CreateInvoice(satoshi, label, description, 3600, nil, "", false)
}

// GetInvoiceByLabel checks if an invoice is paid
func GetInvoiceByLabel(label string) (*glightning.Invoice, error) {
	invoice, err := lightning.GetInvoice(label)
	if err != nil {
		log.Println("Error getting invoice:", err)
		return nil, err
	}
	return invoice, nil
}
