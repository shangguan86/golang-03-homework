package main

import(
	"flag"
	"fmt"
	"time"
)

func Recursion(x int) int {
	if x <= 2 {
		return x
	}
	return Recursion(x-1) + Recursion(x-2)
}

func Iteration(x int) int {
	var s,s1,s2 int
	s1 = 0
	s2 = 1

	if x < 2 {
		return x
	}

	for i := 1; i <= x; i++ {
		s = s1 + s2
		s1 = s2
		s2 = s
	}
	return s
}

func main () {
	var number int
	flag.IntVar(&number, "n", 0, "Successione di Fibonacci")
	flag.Parse()

	if number < 0 {
		fmt.Println("Require a no negtive number,you input is:", number)
		return
	}

	start := time.Now()
	fmt.Printf("Result of fibonacci %d use Recursion is: %d, elapsed %v\n", number, Recursion(number), time.Since(start))

	start = time.Now()
	fmt.Printf("Result of fibonacci %d use Iteration is: %d, elapsed %v\n", number, Iteration(number), time.Since(start))
}
