package main

import (
	"net/http"
	"fmt"
	"log"
	"math/rand"
	"time"
	"math"
)

func generate_randnum() int  {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand_num := rand.Intn(100)
	return rand_num
}

func compute() float64 {
	res := math.Pow(10,2)
	return res
}

func handler_rand(w http.ResponseWriter,r *http.Request)  {
	res := generate_randnum()
	fmt.Fprintf(w,"随机数是:%d" ,res)
}

func handler_pow(w http.ResponseWriter,r *http.Request)  {
	res := compute()
	fmt.Fprintf(w,"10的平方数是:%f" ,res)
}

func main() {
	log.Println("Staring http server on lovalhost:8080")
	http.HandleFunc("/",handler_pow)
	http.HandleFunc("/random",handler_rand)
	http.ListenAndServe(":8080",nil)
}


