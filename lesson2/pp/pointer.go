package main

import (
    "fmt"
)


func main() {
    // 打印int类型数据的内存地址 
    var a = 10
    var p = &a

    //fmt.Println("**p is: ", *(*p))

    fmt.Println("Before assignment, *p: ", *p)
    fmt.Println("Before assignment, p: ", p)
    fmt.Println("Before assignment, &a: ", &a)
    *p = 20
    fmt.Println("After assignment, *p: ", *p)
    fmt.Println("After assignment, p: ", p)
    fmt.Println("After assignment, &a: ", &a)
    fmt.Println(*p)

    var b = 100.0
    fmt.Println(&b)

    var c = true
    fmt.Println(&c)
}
