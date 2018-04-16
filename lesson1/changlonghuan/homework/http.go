package main

import (
	"log"
	"net/http"
	"fmt"
	"math"
)

func handler(w http.ResponseWriter, r *http.Request) {
	res := math.Pow(10,2)
	fmt.Fprint(w, "10 ^ 2 = ",res)
}

func main() {
	log.Println("Staring http server on localhost:8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080",nil)
}
