package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%.1f", math.Pow(10, 2))
}

func main() {
	log.Println("Starting http server on localhost:8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
