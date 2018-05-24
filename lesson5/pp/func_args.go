package main

import "fmt"

func sum(args ...int) (n int) {
	for i := 0; i < len(args); i++ {
		n = n + args[i]
	}
	fmt.Println(n)
	return
}

func min_max(args ...int) (min int, max int) {
	min = args[0]
	max = args[0]
	for i := 0; i < len(args); i++ {
		if args[i] < min {
			min = args[i]
		}
		if args[i] > max {
		}
		max = args[i]
	}
	return
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(min_max(1, 2, 3))
}
