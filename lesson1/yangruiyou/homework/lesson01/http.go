package main

import (
	"net/http"
	"log"
	"fmt"
	"math"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, math.Pow10(2))
}

func main() {
	log.Println("staring http server on localhost 8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
