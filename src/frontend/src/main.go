package main

import (
	"log"
	"net/http"
)

// main starts an http server on the $PORT environment variable.
func main() {
	log.Printf("server starting to listen on port 8080")
	http.HandleFunc("/", index)
    http.ListenAndServe(":8080", nil)
}

// home logs the received request and returns a simple response.
func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}
