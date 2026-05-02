package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain")
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "Hello from KotaCloud!\nRequest: %s\n", r.RequestURI)
    })
    
    fmt.Println("Server starting on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}