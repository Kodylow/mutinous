package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kodylow/mutinous/config"
	"github.com/kodylow/mutinous/db"
	"github.com/kodylow/mutinous/handlers"
	"github.com/kodylow/mutinous/lightning"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

var cfg *config.Config

func ReadmeHandler(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("./README.md")
	if err != nil {
		http.Error(w, "Could not read README.md", http.StatusInternalServerError)
		return
	}

	html := blackfriday.Run(file)
	html = bluemonday.UGCPolicy().SanitizeBytes(html)

	fmt.Fprintf(w, "%s", html)
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config / .env file")
	}

	lightning.InitLightning(cfg.ClnRpcPath)

	_, err = db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/.well-known/lnurlp/{username}", handlers.LNAddressHandler)
	r.HandleFunc("/lnurlp/{username}/callback", handlers.LNURLCallbackHandler)
	r.HandleFunc("/lnurlp/{username}/verify/{label}", handlers.LNURLVerifyHandler)
	r.HandleFunc("/", ReadmeHandler) // added readme handler for index

	log.Println("Starting server on " + cfg.Port + "...")
	http.ListenAndServe(":"+cfg.Port, r)
}
