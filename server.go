package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// PageRequest represents the request payload that the endpoint expects.
type PageRequest struct {
	URI        string `json:"uri"`
	RetryLimit int    `json:"retryLimit"`
}

// PageResponse represents the response payload that the endpoint returns.
type PageResponse struct {
	ID        string `json:"id"`
	URI       string `json:"uri"`
	SourceURI string `json:"sourceUri"`
}

func main() {
	http.HandleFunc("/pagesource", func(w http.ResponseWriter, r *http.Request) {
		// Parse the request payload.
		var req PageRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

	})
	log.Println("Starting server on port 7771...")
	log.Fatal(http.ListenAndServe(":7771", nil))
}
