package main

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
