package main

import (
	"fmt"
	"strings"
)

func toupper(s string) string {
	// 函数是一等公民
	mapper := func(r rune) rune {
		// a: 97
		// A: 65
		// a - A == 32
		return r - 320
	}

	// func(rune) rune
	return strings.Map(mapper, s)
}

func main() {
	fmt.Println(toupper("hello"))
}
