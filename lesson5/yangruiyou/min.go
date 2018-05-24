package main

import "fmt"

func min(first int, rest ...int) int {
	for _, x := range rest {
		if x < first {
			first = x
		}
	}
	return first
}

func main() {
	numbers := []int{7, 6, 2, -1, 7, -3, 9}
	fmt.Println(min(numbers[0], numbers[1:]...))
}
