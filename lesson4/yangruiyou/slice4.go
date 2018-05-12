package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil")
	}

	a := [3]int{1, 2, 3}
	s1 := a[:2]
	s2 := a[:2]
	fmt.Println(s1, s2)
	//slice不可以做比较
	fmt.Println(s1 == s2)
}
