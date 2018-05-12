package main

import (
	"fmt"
)

func ToUpper(s string) string {
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] - 32

	}
	return string(arr)
}

func main() {
	s1 := "hello"
	fmt.Println(ToUpper(s1))
}
