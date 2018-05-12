package main

import (
	"fmt"
)

func Reverse(s string) string {
	strs := []rune(s)
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]

	}
	return string(strs)
}

func main() {
	s := "我爱美丽的中国kkgo"
	fmt.Println(Reverse(s))
	//fmt.Println("kkgo")
}
