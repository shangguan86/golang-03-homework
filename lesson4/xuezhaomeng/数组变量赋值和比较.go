package main

import "fmt"

func main() {

	a1 := [3]int{1, 2, 3}
	fmt.Println(a1)

	var a2 [3]int
	a2 = a1
	fmt.Println(a2)

	fmt.Println(a1 == a2)

	fmt.Println([3]int{1,2} == [3]int{2,2})
	//invalid operation: [3]int literal == [2]int literal (mismatched types [3]int and [2]int)

	for i := 0; i < 3; i++ {
		fmt.Printf("a1[%d]:%d, &a1[%d]: %d, &a1: %d\n", i, a1[i], i, &a1[i], &a1)
	}
	for i := 0; i < 3; i++ {
		fmt.Printf("a2[%d]:%d, &a2[%d]: %d, &a2: %d\n", i, a2[i], i, &a2[i], &a2)
	}

}
