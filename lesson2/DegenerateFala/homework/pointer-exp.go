package main

import (
"fmt"
)

//a = num 是直接给变量重新赋值，p 是指针，指向的a，*p是通过地址修改a的值

func main() {

	// 打印int类型数据的内存地址
	var a = 10
	// p是指向int类型数据的指针
	// *p是a
	// &p是指针p的地址
	var p = &a

	fmt.Println("Before assignment, *p: ", *p) // 打印结果为 10;
	fmt.Println("Before assignment, p: ", p)	  // 打印结果为 a 的内存地址
	fmt.Println("Before assignment, &a: ", &a) // 打印结果为 a 的内存地址

	fmt.Println()

	// 修改指针所指向的数据 #1
	*p = 20
	fmt.Println("After assignment, *p: ", *p) // 打印结果为 20,将 a 内存地址的对应的对象更改为 20
	fmt.Println("After assignment, p: ", p)	 // 打印结果为 a 的内存地址
	fmt.Println("After assignment, &a: ", &a) // 打印结果为 a 的内存地址
	fmt.Println("After assignment, a: ", a)   // 打印结果为 20

	fmt.Println()

	// 直接修改值
	a = 30
	fmt.Println("After assignment, *p: ", *p)  // 打印结果为 30,内存地址依旧不变
	fmt.Println("After assignment, p: ", p)	  // 打印结果为 a 的内存地址
	fmt.Println("After assignment, &a: ", &a)  // 打印结果为 a 的内存地址
	fmt.Println("After assignment, a: ", a)    // 打印结果为 30

}

