package main

import (
	"fmt"
)

func main() {
	a := 15
	fmt.Println(a)
	var p *int
	p = &a
	*p = 20
	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&p)
	var m *int
	n := 50
	m = &n
	p = m
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&p)

}
