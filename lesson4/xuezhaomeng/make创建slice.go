package main

import "fmt"

func main() {
	//	int类型的切片，长度，容量为5
	a := make([]int, 5)
	PrintSlice("a", a)

	// 	int类型切片，长度为0 ，容量为5
	b:=  make([]int ,0 ,5)
	PrintSlice("b", b)


}

func PrintSlice(s string, x []int) {

	fmt.Printf("%s len=%d  cap=%d  %v \n", s, len(x), cap(x), x)

}
