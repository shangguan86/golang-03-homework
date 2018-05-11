package main

import (
	"fmt"
)

func main() {
	s := "hello，golang语言"
	fmt.Println(s)
	runes := []rune(s)
	fmt.Println(runes)
	for _, v := range s {
		fmt.Println(v)
	}
}
