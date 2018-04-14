package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, world!")
}

func main() {
	log.Println("Staring http server on localhost:8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080",nil)
}

