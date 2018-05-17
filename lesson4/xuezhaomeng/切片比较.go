package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	s1 := a[:2]
	s2 := a[:2]
	//fmt.Println(s1 == s2)
	//invalid operation: s1 == s2 (slice can only be compared to nil)
	// 切片不可以做比较,仅可以与nil 做比较
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
}
