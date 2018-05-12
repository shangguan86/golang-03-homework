package main

import "fmt"

func main() {
	var b = false

	c := &b
	*c = true
	fmt.Printf("%v,%v\n", b, *c)
}
