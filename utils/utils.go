package utils

import (
	"crypto/rand"
	"math/big"
	"net/http"
)

// ValidateReplitUser checks if the user exists on replit.com
// This is done by checking if the user's profile page returns a 200 status code or a 404
func ValidateReplitUser(username string) bool {
	resp, err := http.Get("https://replit.com/@" + username)
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	return true
}

// GetMetadata returns the metadata
// LNURLP requires metadata to be used as the desc hash for the invoice
func GetMetadata(username string) string {
	return "[[\"text/identifier\",\"" + username + "@mutinynet.app\"],[\"text/plain\",\"Sats for " + username + "\"]]"
}

// GenerateLabel generates a label for an invoice
func GenerateLabel() (string, error) {
	const n = 32
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}
	return string(ret), nil
}
