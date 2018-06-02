package main

import "fmt"

const (
	a = iota
	b
	_
	c
	d = 5
	e
	f = iota
	g
)

func main() {
	fmt.Println(a, b, c, d, e,f,g)
}
