package main

import (
	"fmt"
	"net/http"
	"strings"
)

func parseAndConstructURL(inputURL string) string {

	parts := strings.Split(inputURL, "/")

	chainAndContract := strings.Split(parts[len(parts)-2], ":")
	chain := chainAndContract[0]
	contractAddress := chainAndContract[1]
	tokenID := parts[len(parts)-1]

	// Construct the new URL
	newURL := fmt.Sprintf("https://surr.app/collect?chain=%s&contract_address=%s&token_id=%s", chain, contractAddress, tokenID)

	return newURL
}

func urlHandler(w http.ResponseWriter, r *http.Request) {
	inputURL := r.URL.Query().Get("url")
	if inputURL == "" {
		http.Error(w, "Missing 'url' query parameter", http.StatusBadRequest)
		return
	}

	newURL := parseAndConstructURL(inputURL)

	fmt.Fprintln(w, newURL)
}

func main() {
	http.HandleFunc("/convert-url", urlHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
