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
	http.HandleFunc("/api/hello", hello)
    http.ListenAndServe(":8080", nil)
}

// home logs the received request and returns a simple response.
func home(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request: %s %s", r.Method, r.URL.Path)
	fmt.Fprintf(w, "Hello, world From a Microservice!")
}

func hello(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprintf(w, "invalid_http_method")
        return
	}
    // Must call ParseForm() before working with data
    r.ParseForm()
    // Log all data. Form is a map[]
    log.Println(r.Form)

    // Print the data back. We can use Form.Get() or Form["name"][0]
    fmt.Fprintf(w, "Hello " + r.Form.Get("name"))
}
