package main

import (
	"fmt"
)

func main() {
	i := 10
	fmt.Println(i)

	s := "hello world"
	fmt.Println(s)

	var vi int
	vi = 20
	fmt.Println(vi)

	var vs string
	vs = "hello golang"
	fmt.Println(vs)

	var (
		v1 int
		v2 string
		v3 float64
		v4 bool
	)
	v1 = 1
	v2 = "abc"
	v3 = 2.2
	v4 = true
	fmt.Printf("%v, %v, %v, %v", v1, v2, v3, v4)
}
