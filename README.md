# Webpage Downloader Server using Go
This is a Go server program that provides an endpoint for downloading a webpage and returning the downloaded file to the client. 

The server accepts a POST request to the `/download` endpoint with a JSON payload containing the URL and retry limit of the webpage to download, retrieves the webpage from the specified URL, downloads the webpage as a local file, and returns a JSON payload with the ID, URI, and source URI of the downloaded file. 

### Features:
* The server also has caching support and will serve the downloaded webpage from the local cache if it has already been requested in the last 24 hours.
* 