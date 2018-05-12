package main

import (
	"flag"
	"fmt"
)

func fib(n int) int {
	if n <= 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	var num = flag.Int("n", 10, "default number for fib")
	flag.Parse()

	fmt.Println("num:", *num)
	fmt.Println("fib result:", fib(*num))
}
