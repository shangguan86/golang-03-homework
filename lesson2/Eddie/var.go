package main

import (
	"fmt"
)

func main() {
	//int
	i := 10
	fmt.Println("short", i)
	var a int
	a = 100
	fmt.Println("long", a)

	//string
	v := "hello world!"
	fmt.Println("short", v)
	var b string
	b = "hello go!"
	fmt.Println("long", b)

	//multi var
	var (
		v1        = 45
		v2 string = "multiVar"
		v3        = 2.0
		v4 int
		v5 = 2.123
		v6 bool
	)
	v6 = false
	fmt.Println(v1, v2)
	fmt.Printf("%0.1f\n", v3)
	fmt.Println(v4)
	fmt.Println(v5)
	fmt.Println(v6)
	fmt.Printf("%v", v5)
}
