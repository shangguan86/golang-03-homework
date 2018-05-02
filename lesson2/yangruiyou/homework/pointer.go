package main

import (
	"fmt"
)

func main() {

	// 打印int类型数据的内存地址
	var a = 10
	// p是指向int类型数据的指针
	// *p是a
	// &p是指针p的地址
	var p = &a
	//指针p取值
	fmt.Println("Before assignment, *p: ", *p)
	//指针p地址
	fmt.Println("Before assignment, p: ", p)
	//变量a地址
	fmt.Println("Before assignment, &a: ", &a)

	fmt.Println()

	// 修改指针所指向的数据 #1
	*p = 20
	//指针p取值
	fmt.Println("After assignment, *p: ", *p)
	//指针p地址
	fmt.Println("After assignment, p: ", p)
	//变量a地址
	fmt.Println("After assignment, &a: ", &a)
	//变量a值
	fmt.Println("After assignment, a: ", a)

	fmt.Println()

	// 直接修改值
	a = 30
	//指针p取值
	fmt.Println("After assignment, *p: ", *p)
	//指针p地址
	fmt.Println("After assignment, p: ", p)
	//变量a地址
	fmt.Println("After assignment, &a: ", &a)
	//变量a值
	fmt.Println("After assignment, a: ", a)

	//var b = 100.0
	//fmt.Println(&b)

	//var c = true
	//fmt.Println(&c)
}
