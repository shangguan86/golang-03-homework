package main

import (
    "fmt"
)

func main() {
    // int
    // i := 10
    var i = 10
    fmt.Println(i)

    // string
    s := "hello, world!"
    fmt.Println(s)

    i1, s1 := 20, "abc"
    fmt.Println(i1)
    fmt.Println(s1)

    var (
        v1 int 
        v2 string
        v3 float64
        v4 bool
    )

    v1 = 1
    v2 = "abc"
    v3 = 2.1
    v4 = true

    fmt.Printf("%d\t%s\t%v\t%t\n", v1, v2, v3, v4)
}
