package main

import (
	"fmt"
	"bufio"
)

func sum(args ...int) int {
	n := 0
	for i := 0; i < len(args); i++ {
		n += args[i]

	}
	return n

}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	fmt.Println(min_max(5, 20, 30, 1))
}

func min_max(args ...int) (min, max int) {
	min = args[0]
	max = args[0]
	for _, x := range args {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}

	}

	return max, min


}
