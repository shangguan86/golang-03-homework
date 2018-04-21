package main

import (
	"fmt"
)

func main() {
	m := 10
	p := &m
	fmt.Println(*p)

	var p *int
	p = &a
	*p = 15
	fmt.Println(*p)

}
