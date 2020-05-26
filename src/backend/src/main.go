package main

import (
	"fmt"
	"log"
	"net/http"
)

// main starts an http server on the $PORT environment variable.
func main() {
	log.Printf("server starting to listen on port 8080")
	http.HandleFunc("/api/hello", hello)
    http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
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

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        fmt.Fprint(w, "custom 404")
    }
}
