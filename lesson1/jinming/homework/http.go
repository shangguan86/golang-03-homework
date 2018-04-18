package main

import (
	"net/http"
	"log"
	"math"
	"strconv"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	a := 10.0
	b := math.Pow(a, 2)
	//s := "10's square equals" + strconv.FormatFloat(utils.FormatFloat64(float64())
	s := "10's square equals" + strconv.FormatFloat(b, 'G', 5, 32)
	fmt.Fprintf(w, s)
}

func main(){
	log.Println("starting http server on localhost:8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080",nil)
}

