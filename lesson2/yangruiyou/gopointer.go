package main

import "fmt"

func main() {
	//申明一个指向int类型数据的指针
	var c = 10
	var p = &c
	fmt.Println(*p)

	a := 100

	fmt.Println(&a)

	b := "hello,51reboot"
	fmt.Println(&b)
}
