package main

import (
    "fmt"
    "strconv"
    "reflect"
)


func main(){
    // 用+号来连接字符串
    s1 := "hello" + "world!"
    //s1 := "中国"
    
    // 随机读取字符
    c := s1[0]
    fmt.Println(string(c))

    // 切片(包含start, 不包含end)
    // 输出: hel
    s2 := s1[0:len(s1)]
    fmt.Println(s2)

    i, _ := strconv.Atoi("-42")
    fmt.Println(i)
    fmt.Println("type of i: ", reflect.TypeOf(i))

    f, _ := strconv.ParseFloat("3.14", 64)
    fmt.Println(f)
    fmt.Println("type of f: ", reflect.TypeOf(f))

    fmt.Println(s1)
}
