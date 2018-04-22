package main

import (
	"flag"
	"fmt"
)

func fib(n int) int {
	var x = 0
	var y = 1
	for i := 0; i < n; i++ {
		x, y = x+y, x
	}
	return x
}

func main() {
	var num int
	flag.IntVar(&num, "n", 0, "the number of fibonacci")
	flag.Parse()

	var f = fib(num)
	fmt.Printf("fibonacci of %d is: %d", num, f)

}
