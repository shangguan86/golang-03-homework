package main

import (
    "fmt"
    "reflect"
)


func main(){
    // xello中国
    s1 := "hello中国"

    //arr := []byte(s1)
    arr := []rune(s1)
    fmt.Println("type of arr: ", reflect.TypeOf(arr))

    arr[0] = 'x'
    arr[5] = '美'

    for k, v := range arr {
        fmt.Println(k, v) 
    }
    
    fmt.Println(s1, string(arr))
}
