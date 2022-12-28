package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	defaultMaxRetries = 10
	cacheExpiration   = 24 * time.Hour
)

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
	// Sanitize the URL by replacing any special characters with underscores
	sanitizedURL := sanitizeURL(url)

	// Check if the webpage has already been requested in the last 24 hours
	filePath := fmt.Sprintf("%s.html", sanitizedURL)
	info, err := os.Stat(filePath)
	if err == nil && time.Since(info.ModTime()) < cacheExpiration {
		fmt.Printf("The following webpage has already been requested in the last 24 hours: %s\n", url)
		// Return the response payload for the cached file
		return &responsePayload{
			ID:        sanitizedURL,
			URI:       url,
			SourceURI: filePath,
		}, nil
	}

	if retryLimit <= 0 || retryLimit > defaultMaxRetries {
		retryLimit = defaultMaxRetries
	}

	for i := 0; i < retryLimit; i++ {
		// Make HTTP GET request to the specified URL
		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		// Create a local file to store the webpage
		file, err := os.Create(filePath)
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
			ID:        sanitizedURL,
			URI:       url,
			SourceURI: filePath,
		}, nil
	}

	// Return an error if the maximum number of retries has been reached
	return nil, fmt.Errorf("Failed to download webpage after %d retries", retryLimit)
}

func sanitizeURL(url string) string {
	// Replace any special characters in the URL with underscores
	url = strings.Replace(url, "http://", "", -1)
	url = strings.Replace(url, "https://", "", -1)
	url = strings.Replace(url, "/", "_", -1)
	url = strings.Replace(url, "?", "_", -1)
	url = strings.Replace(url, "%", "_", -1)
	url = strings.Replace(url, "=", "_", -1)
	url = strings.Replace(url, "&", "_", -1)
	return strings.ReplaceAll(url, "/?%=*:|&<>", "_")
}

func main() {
	http.HandleFunc("/pagesource", func(w http.ResponseWriter, r *http.Request) {
		// Return an error if the request method is not POST
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Return an error for any requests to the root path or any other path
		http.Error(w, "Invalid request path", http.StatusNotFound)
	})

	// Start the server
	log.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)

}
