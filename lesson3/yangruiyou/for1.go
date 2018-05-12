package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("number-%d:%d \n", i, i)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1

	}

	i = 8
	for {
		fmt.Println("c---",i)
		i = i + 1
		if i > 10 {
			break
		}
	}

}
