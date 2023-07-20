package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kodylow/mutinous/handlers"
	"github.com/kodylow/mutinous/lightning"
)

func main() {
	lightning.InitLightning()
	
	r := mux.NewRouter()

	r.HandleFunc("/.well-known/lnurlp/{username}", handlers.LNAddressHandler)
	r.HandleFunc("/lnurlp/{username}/callback", handlers.LNURLCallbackHandler)

	http.ListenAndServe(":8080", r)
}
