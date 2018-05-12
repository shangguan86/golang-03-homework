package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main() {
	a1 := "hello" + "world!"

	r := a1[0]

	fmt.Println(string(r))
	fmt.Println(a1)

	a2 := a1[0:len(a1)]
	println(a2)

	i, _ := strconv.Atoi("-42")
	fmt.Println(i)
	fmt.Println("type of i: ", reflect.TypeOf(i))

	f, _ := strconv.ParseFloat("3.14", 32)
	fmt.Println(f)
	fmt.Println("type of f: ", reflect.TypeOf(f))

}