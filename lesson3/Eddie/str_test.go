package main

import "fmt"

func main() {
	s1 := "hello中国"

	//	arr :=[]byte(s1)
	arr := []rune(s1)
	arr[5] = '美'
	for k, v := range arr {
		fmt.Println(k, v)
	}
	fmt.Println(s1, string(arr))
}
