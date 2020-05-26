package main

import (
	"fmt"
	"log"
	"net/http"
)

// main starts an http server on the $PORT environment variable.
func main() {
	log.Printf("server starting to listen on port 8080")
	http.HandleFunc("/", home)
    http.ListenAndServe(":8080", nil)
}

// home logs the received request and returns a simple response.
func home(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request: %s %s", r.Method, r.URL.Path)
	fmt.Fprintf(w, "Hello, world From a Microservice!")
}
