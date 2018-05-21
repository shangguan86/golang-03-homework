package main

import (
	"fmt"
	"strings"
)

func toUpper(s string) string {
	return strings.Map(func(r rune) rune {
		return r - 32
	}, s)
}

func main() {
	fmt.Println(toUpper("hello"))

}
