package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	s1 := "hello\t" +
		"world!"

	s2 := `
		hello
		world!
		`
	fmt.Println(s1)
	fmt.Println(s2)
	s3 := "hello"
	fmt.Println(string(s3[0:]))
	i, _ := strconv.Atoi("-24")
	fmt.Println(i)
	fmt.Println("type of i:", reflect.TypeOf(i))

	f, _ := strconv.ParseFloat("3.14", 32)
	fmt.Println("type of f:", reflect.TypeOf(f))
}
