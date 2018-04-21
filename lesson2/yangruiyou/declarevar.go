package main

import "fmt"

func main() {
	i := 10
	fmt.Println(i)

	var i1 int = 10
	fmt.Println(i1)

	var i2 = 10
	fmt.Println(i2)

	//strings
	s := "hello,51reboot,kill go."

	fmt.Println(s)

	var s1 string = "hello,51reboot"
	fmt.Println(s1)

	var (
		v1 int
		v2 string
		v3 float64
	)

	v1 = 100
	v2 = "hi"
	v3 = 10.0
	fmt.Printf("%v\t%v\t%v\t", v1, v2, v3)
	
}
