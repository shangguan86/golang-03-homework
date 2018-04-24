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

    fmt.Println("Before assignment, *p: ", *p)	//打印指针p指向的地址的值
    fmt.Println("Before assignment, p: ", p)	//打印指针p的地址
    fmt.Println("Before assignment, &a: ", &a)	//打印变量a的地址

    fmt.Println()

    // 修改指针所指向的数据 #1
    *p = 20			//将指针p指向的地址的值修改成20
    fmt.Println("After assignment, *p: ", *p)	//打印指针p指向的地址的值,20
    fmt.Println("After assignment, p: ", p)	//打印指针p的地址,和a一样
    fmt.Println("After assignment, &a: ", &a)	//打印变量a的地址
    fmt.Println("After assignment, a: ", a)	//打印变量a的值,20

    fmt.Println()

    // 直接修改值
    a = 30				//通过变量赋值,直接修改变量a的值,30
    fmt.Println("After assignment, *p: ", *p) //打印指针p指向的地址的值,30	
    fmt.Println("After assignment, p: ", p)	//打印指针p的地址,和a一样
    fmt.Println("After assignment, &a: ", &a)	//打印变量a的地址
    fmt.Println("After assignment, a: ", a)	//打印变量a的值,30
    
    //var b = 100.0
    //fmt.Println(&b)

    //var c = true
    //fmt.Println(&c)
}
