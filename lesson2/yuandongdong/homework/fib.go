package main

import (
	"flag"
	"fmt"
	"time"
)

//递归方式
func fibRecursion(n int) int {
	if n < 2 {
		return n
	}
	return fibRecursion(n-1) + fibRecursion(n-2)

}

//非递归方式
func fibIteration(n int) int {
	var res, a, b int
	a, b = 0, 1

	if n < 2 {
		return n
	}

	for i := 2; i <= n; i++ {
		res = a + b
		a = b
		b = res
	}
	return res

}

func main() {
	var num int
	flag.IntVar(&num, "n", 0, "the number of fibnacci")
	flag.Parse()

	if num < 0 {
		fmt.Println("Require a no negtive number,you input is:", num)
		return
	}

	start := time.Now()
	fmt.Printf("Result of fibonacci %d use Recursion is: %d, elapsed %v\n", num, fibRecursion(num), time.Since(start))

	start = time.Now()
	fmt.Printf("Result of fibonacci %d use Iteration is: %d, elapsed %v\n", num, fibIteration(num), time.Since(start))
}
