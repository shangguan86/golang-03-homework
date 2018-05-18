package main

import "fmt"

func main() {
	slice := []int{
		1,
		2,
		3,
	}
	array := [2]int{
		6,
		6,
	}
	slice = append(slice, array...)
	fmt.Println(slice)
}
