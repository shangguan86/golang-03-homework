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

	//对变量p初始化并赋值，p保存int变量a的地址
	var p = &a

	//输出p中的指针指向的地址的值
	fmt.Println("Before assignment, *p: ", *p)
	//输出变量p保存的值，即变量a的地址
	fmt.Println("Before assignment, p: ", p)
	//输出变量a的地址
	fmt.Println("Before assignment, &a: ", &a)

	fmt.Println()

	// 修改指针所指向的数据 #1
	//通过指针修改a为20
	*p = 20
	//被修改后，输出为20
	fmt.Println("After assignment, *p: ", *p)
	//变量p保存的地址没有改变
	fmt.Println("After assignment, p: ", p)
	//变量a的地址没有改变
	fmt.Println("After assignment, &a: ", &a)
	//a的值变成20
	fmt.Println("After assignment, a: ", a)

	fmt.Println()

	// 直接修改值
	a = 30
	//a的值变成30
	fmt.Println("After assignment, *p: ", *p)
	//变量p指向的地址不变
	fmt.Println("After assignment, p: ", p)
	//变量a的地址不变
	fmt.Println("After assignment, &a: ", &a)
	//a的值变成30
	fmt.Println("After assignment, a: ", a)

	//var b = 100.0
	//fmt.Println(&b)

	//var c = true
	//fmt.Println(&c)
}
