package main

import (
    "fmt"
)


func main() {

    //  定义变量赋值
    var a = 10
 
    //p 指向变量a的地址
    var p = &a

    fmt.Println("Before assignment, *p: ", *p)  //取p得内容，实际上是取p所指地址的值
    fmt.Println("Before assignment, p: ", p)  //取p， 实际上是取p所指向的地址 
    fmt.Println("Before assignment, &a: ", &a) //取a的地址


    // 修改指针所指向的数据 #1
    *p = 20 //修改指针p的值，实际是修改其所在的地址上的值,a和p公用一个地址
    fmt.Println("After assignment, *p: ", *p) //取p得内容，实际上是取p所指地址的值
    fmt.Println("After assignment, p: ", p)
    fmt.Println("After assignment, &a: ", &a)
    fmt.Println("After assignment, a: ", a)


    // 直接修改值
    a = 30 //修改指针a的值，实际是修改其所在的地址上的值,p和a公用一个地址
    fmt.Println("After assignment, *p: ", *p) 
    fmt.Println("After assignment, p: ", p)
    fmt.Println("After assignment, &a: ", &a)
    fmt.Println("After assignment, a: ", a)
    
    //var b = 100.0
    //fmt.Println(&b)

    //var c = true
    //fmt.Println(&c)
}
