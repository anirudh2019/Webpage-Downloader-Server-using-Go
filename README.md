# Webpage Downloader Server using Go
This is a Go server program that provides an endpoint for downloading a webpage and returning the downloaded file to the client. 

The server accepts a POST request to the `/download` endpoint with a JSON payload containing the URL and retry limit of the webpage to download, which then retrieves the webpage from the specified URL, downloads the webpage as a local file, and returns a JSON payload with the ID, URI, and source URI of the downloaded file. 

### Features:
* The server retries maximum upto 10 times or the specified retry limit, whichever is lower, before either successfully downloading the webpage or marking the page as a failure.

* The server also has caching support and will serve the downloaded webpage from the local cache if it has already been requested in the last 24 hours.

* Responding with HTTP errors if there is an issue with the request or the server 

### Installation:
To install the Webpage Downloader Server, you will need to have Go installed on your machine. You can download Go from the [official website](https://go.dev/dl/) and follow the installation instructions.

Once Go is installed, you can download the source code for the Webpage Downloader Server using the following command:
```go
go get github.com/anirudh2019/Webpage-Downloader-Server-using-Go
```
This will download the source code for the Webpage Downloader Server to your local machine.