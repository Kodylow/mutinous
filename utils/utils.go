package utils

import (
	"net/http"
)

// ValidateReplitUser checks if the user exists on replit.com
func ValidateReplitUser(username string) bool {
	resp, err := http.Get("https://replit.com/@" + username)
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	return true
}
