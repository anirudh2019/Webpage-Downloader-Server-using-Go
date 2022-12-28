# Webpage Downloader Server using Go
This is a Go server program that provides an endpoint for downloading a webpage and returning the downloaded file to the client. 

The server accepts a POST request to the `/pagesource` endpoint with a JSON payload containing the URL and retry limit of the webpage to download, which then retrieves the webpage from the specified URL, downloads the webpage as a local file, and returns a JSON payload with the ID, URI, and source URI of the downloaded file. 

Table of contents
=================
<!--ts-->
* [Features](#features)
* [Installation](#installation)
* [Usage](#usage)
* [Assumptions](#assumptions)
<!--te-->

## Features:
* The server retries maximum upto 10 times or the specified retry limit, whichever is lower, before either successfully downloading the webpage or marking the page as a failure.

* The server also has caching support and will serve the downloaded webpage from the local cache if it has already been requested in the last 24 hours.

* Responding with HTTP errors if there is an issue with the request or the server 

## Installation:
To install the Webpage Downloader Server, you will need to have Go installed on your machine. You can download Go from the [official website](https://go.dev/dl/) and follow the installation instructions.

Once Go is installed, you can download the source code for the Webpage Downloader Server using the following command:
```go
go get github.com/anirudh2019/Webpage-Downloader-Server-using-Go
```
This will download the source code for the Webpage Downloader Server to your local machine.

## Usage:
To use the Webpage Downloader Server, you will need to start the server by running the following command from the root directory of the project:
```go
go run server.go
```
This will start the server and listen for incoming HTTP requests on port 8080.

To download a webpage, you can send a POST request to the `/pagesource` endpoint with a JSON payload containing the URL and retry limit of the webpage to download. For example, you can use the `curl` command to send a request like this:

```
curl --location --request POST http://localhost:8080/pagesource --header "Content-Type: application/json" --data-raw "{\"url\": \"https://www.example.com\", \"retry_limit\": 5}"
```

This will send a POST request to the `/pagesource` endpoint with a JSON payload containing the URL "https://www.example.com" and a retry limit of 5. The server will retrieve the webpage from the specified URL and download it as a local file, then return a JSON payload with the ID, URI, and source URI of the downloaded file.

The JSON payload returned by the server will have the following structure:

```
{
	"id": "sanitized_url",
	"uri": "sanitized_url.html",
	"source_uri": "original_url"
}
```
Where `sanitized_url` is the URL of the webpage with any special characters replaced with underscores, and `original_url` is the original URL of the webpage specified in the request payload.

## Assumptions:
* The server is designed to accept POST requests to the `/pagesource` endpoint. It does not accept GET requests. It also does not accept requests to any path other than `/pagesource`.
* The server is running on localhost at port 8080
* The URL in the request payload is a valid URL
* The retry limit in the request payload is a positive integer
* The server has permission to create and modify files in the current working directory
