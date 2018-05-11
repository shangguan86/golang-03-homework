package main

import "fmt"

func main() {
	var b = false
	var a *bool = &b
	*a = true
	fmt.Printf("b value:%v", b)
}
