package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
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
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Parse the request payload.
		var req PageRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Fetch the webpage and download it as a file.
		var err error
		var resp *http.Response
		for i := 0; i < req.RetryLimit; i++ {
			resp, err = http.Get(req.URI)
			if err == nil {
				break
			}
			time.Sleep(time.Second)
		}
		defer resp.Body.Close()

		// Generate a unique ID for the file.
		id := fmt.Sprintf("%x", time.Now().UnixNano())

		// Save the file to the local file system.
		file, err := os.Create(fmt.Sprintf("/files/%s.html", id))
		defer file.Close()

		// Return the response payload.
		res := PageResponse{
			ID:        id,
			URI:       req.URI,
			SourceURI: fmt.Sprintf("/files/%s.html", id),
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	})
	log.Println("Starting server on port 7771...")
	log.Fatal(http.ListenAndServe(":7771", nil))
}
