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
	//p是变量a的指针

	fmt.Println("Before assignment, *p: ", *p)
	//＊p 是指针，指向&a的内存地址所以*(&a)实际取的是a的值。
	fmt.Println("Before assignment, p: ", p)
	// p是变量，因为p＝&a，&a是变量a的内存地址，所以p实际为a的内存地址。
	fmt.Println("Before assignment, &a: ", &a)
	// a是变量，&a为变量的值10的内存地址。

	fmt.Println()

	// 修改指针所指向的数据 #1
	*p = 20
	fmt.Println("After assignment, *p: ", *p)
	//因为p＝&a，＊p ＝＊（&a），但是＊p＝20，重新赋值＊p，所以＊p是一个指向值20的指针。
	fmt.Println("After assignment, p: ", p)
	//P仍然是一个指向a的内存地址的变量，不受赋值影响。
	fmt.Println("After assignment, &a: ", &a)
	//a不受影响，&a为变量10的内存地址。
	fmt.Println("After assignment, a: ", a)
	// a变量的值10

	fmt.Println()

	// 直接修改值
	a = 30
	fmt.Println("After assignment, *p: ", *p)
	// *p是指针指向&a，&a为a的内存地址所以＊p为＊（&a）=30
	fmt.Println("After assignment, p: ", p)
	// ?
	fmt.Println("After assignment, &a: ", &a)
	// ?
	fmt.Println("After assignment, a: ", a)
	// a 重新赋值30
	fmt.Println(*(&a))

	//var b = 100.0
	//fmt.Println(&b)

	//var c = true
	//fmt.Println(&c)
}
