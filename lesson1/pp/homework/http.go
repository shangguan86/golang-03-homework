package main

import (
    "fmt"
    "net/http"
    "log"
    "math"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%f", math.Pow10(2))
}

func main() {
    log.Println("Starting http server on localhost:8080")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
