package main

import "fmt"

func main(){
	var a =10
	var p =&a	//定义一个int类型的指针p,指向int类型的变量a
	fmt.Println("变量a的地址:",&a)	//变量a的地址: {地址a}
	fmt.Println("变量a的值:",a)	//变量a的值: 10
	fmt.Println("p 是指针的指向的地址:",p)	//p 是指针的指向的地址: {地址a}
	fmt.Println("&p 是指针p的地址:",&p)	//&p 是指针p的地址: {地址p}
	fmt.Println("*p 是指针p指向的内容:",*p)	//*p 是指针p指向的内容: 10

	a = 20
	fmt.Println("变量a的地址:",&a)	//变量a的地址: {地址a}
	fmt.Println("变量a的值:",a)	//变量a的值: 20
	fmt.Println("p 是指针指向的地址:",p)	//p 是指针指向的地址: {地址a}
	fmt.Println("&p 是指针p的地址:",&p)	//&p 是指针p的地址: {地址p}
	fmt.Println("*p 是指针p指向的内容:",*p)	//*p 是指针p指向的内容: 20

	*p = 10		//将指针指向的变量的值从20修改到10
	fmt.Println("变量a的地址:",&a)	//变量a的地址: {地址a}
	fmt.Println("变量a的值:",a)	//变量a的值: 10
	fmt.Println("p 是指针的指向的地址:",p)	//p 是指针的指向的地址: {地址a}
	fmt.Println("&p 是指针p的地址:",&p)	//&p 是指针p的地址: {地址p}
	fmt.Println("*p 是指针p指向的内容:",*p)	//*p 是指针p指向的内容: 10
}
