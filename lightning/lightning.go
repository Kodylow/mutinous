package lightning

import (
	"github.com/niftynei/glightning/glightning"
	"log"
	"os/exec"
)

// Custom wrapper to limit the scope of the lightning client
var lightning *glightning.Lightning

// InitLightning initializes the lightning client
func InitLightning(rpcPath string) {
	cmd := exec.Command("lightningd", "--lightning-dir=/home/runner/mutinous/.lightning/", "--signet", "--disable-plugin", "bcli")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("CLN daemon started with process id %d...\n", cmd.Process.Pid)

	lightning = glightning.NewLightning()
	log.Println("RPC path:", rpcPath)
	lightning.StartUp(rpcPath, "lightning-rpc")

	log.Printf("Lightning client initialized...\n")
}

// CreateInvoice creates an invoice
func CreateInvoice(satoshi uint64, label string, description string) (*glightning.Invoice, error) {
	return lightning.CreateInvoice(satoshi, label, description, 3600, nil, "", false)
}

// IsInvoicePaid checks if an invoice is paid
func GetInvoiceByLabel(label string) (*glightning.Invoice, error) {
	return lightning.GetInvoice(label)
}
