package main

import (
    "fmt"
)

/*
其实可以比喻为去商场存包:
	商场的柜子就是内存
	号牌就是指针
	包就是向内存里面值
	可以通过号牌来更换存的包
 */
 
func main() {

    // 打印int类型数据的内存地址
    var a = 10
    // p是指向int类型数据的指针
    // *p是a
    // &p是指针p的地址
    //声明p并且指向a的内存地址
    var p = &a
    //打印p的值(其实是a的值)
    fmt.Println("Before assignment, *p: ", *p)
    //打印p(其实是a的内存地址)
    fmt.Println("Before assignment, p: ", p)
    //打印a的内存地址
    fmt.Println("Before assignment, &a: ", &a)

    fmt.Println()

    // 修改指针所指向的数据 #1
    //通过内存地址修改a的值为20
    *p = 20
    fmt.Println("After assignment, *p: ", *p)
    fmt.Println("After assignment, p: ", p)
    fmt.Println("After assignment, &a: ", &a)
    fmt.Println("After assignment, a: ", a)

    fmt.Println()

    // 直接修改值
    a = 30
    fmt.Println("After assignment, *p: ", *p)
    fmt.Println("After assignment, p: ", p)
    fmt.Println("After assignment, &a: ", &a)
    fmt.Println("After assignment, a: ", a)

    //var b = 100.0
    //fmt.Println(&b)

    //var c = true
    //fmt.Println(&c)
}
