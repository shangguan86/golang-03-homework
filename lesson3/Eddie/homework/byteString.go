package main

import (
	"fmt"
)

func main() {
	str2 := "hello"
	data2 := []byte(str2)
	fmt.Println(data2)
	str2 = string(data2[:])
	fmt.Println(str2)
	str3 := "中国"
	data3 := []rune(str3)
	fmt.Println(data3)
}
