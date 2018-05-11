package main

import "fmt"

func main () {
	var a1 = 0.1110
	var a2 = 0.1111
	var a = 3.0/2

	var equal = (a1 - a2) < 0.01

	fmt.Println(equal)


	fmt.Println(a)
}
