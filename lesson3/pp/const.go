package main

import (
	"fmt"
)

const (
	A = 11 * iota
	B
	C
	D
)

var (
	p = 3.15
	b = 3.25
)

func main() {
	//PI = 3.15
	fmt.Println(A, B, C, D)
}
