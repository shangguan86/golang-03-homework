package main

import "fmt"

func reverse(s string) string {
	sli := []rune(s)
	//fmt.Println(len(sli)/2)
	for i := 0; i < len(sli)/2; i++ {
	//	fmt.Println(i)
		sli[i], sli[len(sli)-i-1] = sli[len(sli)-i-1], sli[i]

	}
	return string(sli)

}

func main() {
	char := "hello world"
	r_char := reverse(char)
	fmt.Printf("原来字符串: %s\n", char)
	fmt.Printf("反转字符串: %s\n", r_char)
}
