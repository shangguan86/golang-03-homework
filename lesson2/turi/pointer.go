package main

import "fmt"

func main() {
	a := 10
	p := &a
	fmt.Println(*p)

	b := 100.0
	fmt.Println(&b)

	c := true
	fmt.Println(&c)

}
