package main

import "fmt"

//多值循环原地反转
func reverse1(s string) string {
	s1 := []rune(s)
	for i, j := 0, len(s1)-1; i < j; i, j = i+1, j-1 {
		s1[i], s1[j] = s1[j], s1[i]
	}
	return string(s1)
}

//循环原地反转
func reverse2(s string) string {
	s1 := []rune(s)
	for i := 0; i < len(s1)/2; i++ {
		s1[i], s1[len(s)-i-1] = s1[len(s)-i-1], s1[i]
	}
	return string(s1)
}

//copy
func reverse3(s string) string {
	runes1 := []rune(s)
	runes2 := make([]rune, len(runes1))
	for i := 0; i < len(runes1); i++ {
		runes2[i] = runes1[len(runes1)-i-1]
	}
	return string(runes2)
}

func main() {
	x := "hello"
	fmt.Println(reverse1(x))
	fmt.Println(reverse2(x))
	fmt.Println(reverse3(x))
}
