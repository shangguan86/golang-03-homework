// package
package main

import (
	"fmt"
	"net/http"
	"log"
	"math"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, math.Pow(10, 2))
}

func main() {
	log.Println("starting listen on localhost:8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}