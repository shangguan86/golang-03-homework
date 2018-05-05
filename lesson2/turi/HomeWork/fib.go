package main

import (
	"fmt"
	"time"
)

func fib1(f1 int) int {
	if f1 < 2 {
		return f1
	}

	return fib1(f1-1) + fib1(f1-2)
}

func fib2(f2 int) int {
	a := 0
	b := 1
	for i := 0; i < f2; i++ {
		a, b = b, a+b
	}
	return a
}

func main() {
	t0 := time.Now()
	fmt.Println(fib1(10))
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

	t2 := time.Now()
	fmt.Println(fib2(10))
	t3 := time.Now()
	fmt.Printf("The call took %v to run.\n", t3.Sub(t2))

}
