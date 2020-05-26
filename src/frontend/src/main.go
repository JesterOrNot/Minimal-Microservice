package main

import (
	"fmt"
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
	if r.URL.Path != "/" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
	http.ServeFile(w, r, "./static/index.html")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        fmt.Fprint(w, "custom 404")
    }
}
