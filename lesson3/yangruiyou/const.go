package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main() {

	s1 := "hello" + "world"

	s2 := `helll-c
hello
hello
hello
`
	fmt.Println(s1)
	fmt.Println(s2)

	i, _ := strconv.Atoi("-42")
	fmt.Println(i)
	fmt.Println("type of i:", reflect.TypeOf(i))
}
