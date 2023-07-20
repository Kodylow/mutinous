package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kodylow/mutinous/config"
	"github.com/kodylow/mutinous/handlers"
	"github.com/kodylow/mutinous/lightning"
)

var cfg *config.Config



func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	lightning.InitLightning(cfg.ClnRpcPath)
	
	r := mux.NewRouter()

	r.HandleFunc("/.well-known/lnurlp/{username}", handlers.LNAddressHandler)
	r.HandleFunc("/lnurlp/{username}/callback", handlers.LNURLCallbackHandler)

	http.ListenAndServe(":" + cfg.Port, r)
}
