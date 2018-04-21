package main

import "fmt"

func main() {
	//申明一个指向int类型数据的指针
	var c = 10
	var p = &c

	fmt.Println("before assign ,p:", p)   //取指针p的地址
	fmt.Println("before assign ,*p:", *p) //取指针p指向的值
	fmt.Println("before assign ,&p:", &p) //?

	//*p = 100

	//a := 100

	//fmt.Println(&a)

	//b := "hello,51reboot"
	//fmt.Println(&b)
}
