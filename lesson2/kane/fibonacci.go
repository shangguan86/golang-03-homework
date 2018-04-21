package main

import (
	"flag"
	"fmt"
)

func main() {
	var num int
	flag.IntVar(&num, "n", 0, "The steps")
	flag.Parse()

	if num <= 0 {
		fmt.Println("Error")
	}

	var ways = fib(num)
	fmt.Println("递归的结果", ways)

	var i int
	var first int = 1
	var second int = 2
	for i = 1; i < num; i++ {
		first, second = step(first, second)
	}
	fmt.Println("循环的结果", first)
}
func step(first int, second int) (int, int) {
	return second, second + first
}

func fib(n int) int {
	if n <= 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
