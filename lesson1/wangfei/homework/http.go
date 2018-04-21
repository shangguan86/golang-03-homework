package main

import (
	"net/http"
	"fmt"
	"log"
	"math/rand"
	"time"
	"math"
)


func handler_rand(w http.ResponseWriter,r *http.Request)  {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := rand.Intn(100)
	fmt.Fprintf(w,"随机数是:%d" ,res)
}

func handler_pow(w http.ResponseWriter,r *http.Request)  {
	res := math.Pow(10,2)
	fmt.Fprintf(w,"10的平方数是:%f" ,res)
}

func main() {
	log.Println("Staring http server on lovalhost:8080")
	http.HandleFunc("/",handler_pow)
	http.HandleFunc("/random",handler_rand)
	http.ListenAndServe(":8080",nil)
}


