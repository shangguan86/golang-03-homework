package main

import (
	"fmt"
)

func reverse(str string) string {
	if len(str) < 1 {
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
	str := "vim-go"
	fmt.Println("Source string: ", str)
	fmt.Println("Reversed string: ", reverse("vim-go"))
}
