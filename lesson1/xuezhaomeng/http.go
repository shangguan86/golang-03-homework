package main

import (
	"fmt"
	"log"
	"net/http"
)

func hander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")

}

func main() {

	log.Println("Starting  http server  on localhost:8080")
	http.HandleFunc("/", hander)
	http.ListenAndServe(":8080", nil)
}
