package main

import (
	"fmt"
	"net/http"
	"log"
	"math"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	v := math.Pow(10, 2)
	fmt.Fprintf(w, "10 * 10 = "+strconv.FormatFloat(v,'G',5,32))
}

func main() {
	log.Println("Starting http server on localhost:8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
