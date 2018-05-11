package main

import "fmt"

func main() {

	str := "hello"
	arr := []rune(str)
	var s string
	for i := len(arr) - 1; i <= len(arr)-1; i-- {
		//fmt.Println(i)
		//fmt.Println(string(arr[i]))
		s = s + string(arr[i])
		if i == 0 {
			break
		}

	}
	fmt.Println(s)

}
