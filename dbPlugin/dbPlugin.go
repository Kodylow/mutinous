package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/elementsproject/glightning/glightning"
)

var plugin *glightning.Plugin

func main() {
	plugin = glightning.NewPlugin(onInit)
	registerHooks(plugin)

	err := plugin.Start(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

func onInit(plugin *glightning.Plugin, options map[string]glightning.Option, config *glightning.Config) {
	log.Printf("successfully init'd! %s\n", config.RpcFile)
}

func OnInvoicePaid(paymentEvent *glightning.InvoicePaymentEvent) (*glightning.InvoicePaymentResponse, error) {
	payment := paymentEvent.Payment
	log.Printf("invoice paid for amount %s with preimage %s", payment.MilliSatoshis, payment.PreImage)

	// Notify the main app
	jsonData, err := json.Marshal(payment)
	if err != nil {
		log.Println("Failed to marshal payment:", err)
		return nil, err
	}

	resp, err := http.Post("http://localhost:8080/invoicePaid", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Failed to notify main app:", err)
		return nil, err
	}
	defer resp.Body.Close()

	log.Println("Successfully notified main app")
	return nil, nil
}

func registerHooks(p *glightning.Plugin) {
	p.RegisterHooks(&glightning.Hooks{
		InvoicePayment: OnInvoicePaid,
	})
}
