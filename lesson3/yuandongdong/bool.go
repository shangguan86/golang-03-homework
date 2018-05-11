package main

import "fmt"

func main() {
	var b = false
	a := &b
	*a = true
	fmt.Printf("%v\n", b)
}
