package main

import "fmt"

func reverse_string(str string) {
	arr := []rune(str)
	var arr_reserve []rune
	for i := 1; i < len(arr)+1; i++ {
		arr_reserve = append(arr_reserve, arr[len(arr)-i])
	}
	fmt.Printf("原字符串:%v\n", str)
	fmt.Printf("反转后字符串:%v\n", string(arr_reserve))
}


func main() {
	str := "hello world!"
	reverse_string(str)
}
