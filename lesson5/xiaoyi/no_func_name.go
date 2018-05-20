package main

import (
	"strings"
	"fmt"
)

//匿名函数1
func toupper(s string) string {
	mapper := func(r rune) rune {
		return r - 32
	}
	return strings.Map(mapper,s)
}

//闭包函数
func squares() func() int {
	var x int
	return func() int {
		x++
		return x*x
	}
}

func main() {
	fmt.Println(toupper("hello"))

	var f func() int = squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
