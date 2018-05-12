package main

import (
	"fmt"
)

func main() {
	var inputCout int
	fmt.Scanln(&inputCout)

	var arr = make([]int, inputCout, inputCout)
	for i := 0; i < inputCout; i++ {
		fmt.Scanln(&arr[i])
	}
	for i := inputCout - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}
}
