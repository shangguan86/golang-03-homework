package main

import (
	"fmt"
	"os"
)

func reverse(str string) string {
	if len(str) < 2 {
		return str
	}

	var res, strRune []rune
	strRune = []rune(str)

	for i := len(strRune) - 1; i >= 0; i-- {
		res = append(res, strRune[i])
	}

	return string(res)

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("请输入一个字符串，你输入的参数是：", os.Args[1:])
		return
	}

	str := os.Args[1]
	fmt.Println("Source string: ", str)
	fmt.Println("Reversed string: ", reverse(str))
}
