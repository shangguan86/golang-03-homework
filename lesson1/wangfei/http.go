package main

import (
	"net/http"
	"fmt"
	"log"
)

func handler(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"Hello World")
}

func main() {
	log.Println("Staring http server on lovalhost:8080")
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080",nil)
}

