package main

import "fmt"

func main() {
	var b = false
	var p = &b
	*p = 2 > 1
	fmt.Printf("%v\n", b)
}
