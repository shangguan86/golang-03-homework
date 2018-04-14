package main

import (
	"net/http"
	"fmt"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "hello, world!")
}

func main(){
	log.Println("starting http server on localhost:8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080",nil)
}
