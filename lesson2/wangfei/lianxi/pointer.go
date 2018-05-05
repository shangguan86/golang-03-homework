package main

import "fmt"

func main(){
	//声明一个指向int类型数据的指针

	//int类型数据的内存地址
	//var a =10
	//fmt.Println(&a)
	//
	//var s = "hell world"
	//fmt.Println(&s)
	//
	//var f = 2.0
	//fmt.Println(&f)

	//var b bool
	//b = false
	//fmt.Println(&b)
	//fmt.Println(*&b)

	//var b =100.0
	//fmt.Println(&b)

	var a =10
	var p =&a
	fmt.Println("变量a的内存地址",&a)
	fmt.Println("变量a的值:",a)

	fmt.Println("p 是指针p的指向的地址:",p)
	fmt.Println("&p 是指针p的地址",&p)
	fmt.Println("*p 是指针p指向的地址所对应的值",*p)

	a = 20
	fmt.Println("p 是指针p的指向的地址:",p)
	fmt.Println("&p 是指针p的地址",&p)
	fmt.Println("*p 是指针p指向的地址所对应的值",*p)

	*p = 10
	fmt.Println("p 是指针p的指向的地址:",p)
	fmt.Println("&p 是指针p的地址",&p)
	fmt.Println("*p 是指针p指向的地址所对应的值",*p)

}
