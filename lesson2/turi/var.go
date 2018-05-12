package main

import (
	"fmt"
)

func main() {
	// int
	// i := 10
	// var i int = 10
	// fmt.Println(i)

	// string
	// s := "hello world!"
	// fmt.Println(s)

	var (
		v1 int
		v2 float64
		v3 string
		v4 bool
	)
	v1 = 1
	v2 = 2.1
	v3 = "hello"
	v4 = true
	fmt.Printf("%v,%v,%v,%v\n", v1, v2, v3, v4)
}
