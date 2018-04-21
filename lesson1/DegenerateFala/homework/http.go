package main

import (
    "fmt"
    "net/http"
    "log"
	"math"
)

func handler(w http.ResponseWriter, r *http.Request) {
	const a  = 10
	//也可以指定变量类型
	//var a float64
	//a = 10
	b := math.Pow(a ,2)
	fmt.Fprintln(w, b)
}

func main() {
    log.Println("Starting http server on localhost:8080")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
