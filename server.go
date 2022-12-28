package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const defaultMaxRetries = 10

type requestPayload struct {
	URL        string `json:"url"`
	RetryLimit int    `json:"retry_limit"`
}

type responsePayload struct {
	ID        string `json:"id"`
	URI       string `json:"uri"`
	SourceURI string `json:"source_uri"`
}

func downloadPage(url string, retryLimit int) (*responsePayload, error) {
	if retryLimit <= 0 {
		retryLimit = defaultMaxRetries
	}

	for i := 0; i < retryLimit; i++ {
		// Make HTTP GET request to the specified URL
		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		// Generate a unique ID for the downloaded file
		id := fmt.Sprintf("%d", time.Now().UnixNano())

		// Create a local file to store the webpage
		file, err := os.Create(fmt.Sprintf("%s.html", id))
		if err != nil {
			return nil, err
		}
		defer file.Close()

		// Copy the response body to the local file
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			continue
		}

		// Return the response payload
		return &responsePayload{
			ID:        id,
			URI:       file.Name(),
			SourceURI: url,
		}, nil
	}

	// Return an error if the maximum number of retries has been reached
	return nil, fmt.Errorf("Failed to download webpage after %d retries", retryLimit)
}

func main() {
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		// Parse the request payload
		var payload requestPayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Call the downloadPage function to download the webpage
		resp, err := downloadPage(payload.URL, payload.RetryLimit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return the response payload as JSON
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, "Error encoding response payload", http.StatusInternalServerError)
			return
		}
	})
	// Start the server
	http.ListenAndServe(":8080", nil)
}
