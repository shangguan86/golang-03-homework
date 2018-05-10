package main

import "fmt"

//定义一个常量，注意常量不可变
const (
	PI = 3.15
	A = iota
	B
	C
	D
)

func main() {
	//PI = 3.15  常量不可更改
	fmt.Println(PI)
	fmt.Println(A,B,C,D)

}
