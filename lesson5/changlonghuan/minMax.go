package main

import (
	"fmt"
)

func minmax(args ...int) (min, max int) {
	min = args[0]
	max = args[0]
	for i := 0; i < len(args); i++ {
		a := args[i]
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
	}
	return
}
func main() {
	fmt.Println(minmax(1, 2, 3, 4))
}
