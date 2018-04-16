package main

import (
	"fmt"
	"net/http"
	"log"
	"math"
)

func handler(w http.ResponseWriter, r *http.Request) {
	s := math.Pow(10,2)
	fmt.Fprintln(w,s)
}

func main() {
	log.Println("Starting http server on localhost:8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
