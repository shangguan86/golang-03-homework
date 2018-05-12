package main

import "fmt"

func main() {

	var b = false
	fmt.Printf("%v\n", b)
	p := &b
	*p = true
	fmt.Printf("%v\n", b)

}
