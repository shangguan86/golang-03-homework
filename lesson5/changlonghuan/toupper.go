package main

import (
	"fmt"
	"strings"
)

func toupper(s string) string {
	//没有函数名,就是匿名函数
	mapFunc := func(r rune) rune {
		return r - ('A' - 'a')
	}
	return strings.Map(mapFunc, s)
}
func main() {
	fmt.Println(toupper("hello"))
}