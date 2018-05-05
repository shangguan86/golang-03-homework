package main

import (
	"fmt"
)

func main() {

	//声明int类型变量a并赋值
	var a = 10
	//声明int指针类型变量p并赋值变量a的内存地址
	var p = &a

	//打印p指针指向内存地址的数据 10 a的值
	fmt.Println("Before assignment, *p: ", *p)
	//打印p指针变量 a的内存地址
	fmt.Println("Before assignment, p: ", p)
	//打印a的内存地址
	fmt.Println("Before assignment, &a: ", &a)

	fmt.Println()

	*p = 20
	//打印p指针指向内存地址的数据 20
	fmt.Println("After assignment, *p: ", *p)
	//打印p指针变量 a的内存地址
	fmt.Println("After assignment, p: ", p)
	//打印a的内存地址
	fmt.Println("After assignment, &a: ", &a)
	//打印a变量 20
	fmt.Println("After assignment, a: ", a)

	fmt.Println()

	a = 30
	//打印p指针指向内存地址的数据 30
	fmt.Println("After assignment, *p: ", *p)
	//打印p指针变量 a的内存地址
	fmt.Println("After assignment, p: ", p)
	//打印a的内存地址
	fmt.Println("After assignment, &a: ", &a)
	//打印a变量 20
	fmt.Println("After assignment, a: ", a)

}
