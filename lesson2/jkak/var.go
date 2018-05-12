package main

import "fmt"

var (
	i = 10
	s = "hello golang"
)

func main() {
	fmt.Println(i)
	fmt.Println(s)

	v1 := 1
	v2 := "abc"
	v3 := 2.1
	fmt.Printf("%v\t%v\t%v\n", v1, v2, v3)
}
