package main

import "fmt"

func triple(x int) (result int) {
	defer func() {
		result += x
		fmt.Printf("double(%d)=%d\n", x, result)
	}()

	return x + x
}

func main() {
	fmt.Println(triple(2))
}
