package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func hander(w http.ResponseWriter, r *http.Request) {
	value := math.Pow(10, 2)
	fmt.Fprintln(w, value)
	fmt.Fprintf(w, "%v\n", value)

}

func main() {

	log.Println("Starting  http server  on localhost:8080")
	http.HandleFunc("/", hander)
	http.ListenAndServe(":8080", nil)
}
