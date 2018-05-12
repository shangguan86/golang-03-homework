package main

import (
	"fmt"
)

func main() {
	s := "hello，golang语言"
	fmt.Println(s)
	runes := []rune(s)
	fmt.Println(runes)
	n := len(runes)
	fmt.Println(n)
	fmt.Println(runes[13])
	// var arr = make([]rune, n, n)
	// for k, _ := range arr {
	for i := 0; i <= n/2; i++ {
		tmp := runes[i]
		runes[i] = runes[n-i-1]
		runes[n-1-i] = tmp
	}

	// }

	fmt.Println(string(runes))
	// for _, v := range s {
	// fmt.Println(v)
	// }
}
