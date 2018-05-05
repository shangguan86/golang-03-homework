package main

import "fmt"

func main() {

	// for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	// while
	i := 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println()

	// while true
	i = 8
	for {
		fmt.Println(i)
		i = i + 1
		if i > 10 {
			break
		}
	}

}
