package main

import (
	"fmt"
)

func toUpper(s string) string {
	arr := []rune(s)
	for k, v := range arr {
		if arr[k] >= 97 && arr[k] <= 122 {
			arr[k] = v - 32
			fmt.Println(arr)
		}
	}
	return string(arr)
}
func main() {
	s1 := "hello中国"
	fmt.Println(toUpper(s1))
}
