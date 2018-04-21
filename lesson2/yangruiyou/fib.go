package main

import (
	"fmt"
	"flag"
)

func fib(n int) int {

	if n <= 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}

}

func main() {

	var num int
	flag.IntVar(&num, "n", 0, "number of steps")
	flag.Parse()

	fmt.Printf("fib(%v) is :%v \n", num, fib(num))

}
