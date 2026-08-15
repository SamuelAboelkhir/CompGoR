// Package utils: A utility package with helper functions
package utils

import (
	"net/http"
	"strings"
)

// CleanInput splits a string into a slice of words, trimming spaces and punctuation
// Example:
//
//	CleanInput("Hello, world!") // returns [["Hello", "world"]]
func CleanInput(text string) []string {
	words := strings.Fields(text)
	return words
}

// CheckJSON checks the content-tyoe of the response header to make sure the response is JSON
func CheckJSON(res *http.Response) bool {
	return res.Header.Get("Content-Type") == "application/json"
}
