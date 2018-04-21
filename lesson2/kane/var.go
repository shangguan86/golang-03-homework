package main

import "fmt"

func main() {
	//int
	i:=10
	fmt.Println(i)

	var i1 int64 = 10

	fmt.Println(i1)

    // string
	s := "hello, world!"
	fmt.Println(s)

	//var s1 string  = "hello, world!"

	var (
		v1 int
		v2 float64
		v3 string
		v4 bool
	)

    v4 = true
    v3 = "sdfaasdfs"

	fmt.Printf("%v\t%f\t%v\t%s\t",v1,v2,v4,v3)

	fmt.Println(v1,v2,"["+v3+"]",v4)
}
