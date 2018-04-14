package main

import (
    "fmt"
    "os"
    "strings"
)

func test1()  {
    name := [...]string{"a","b"}
    fmt.Println(strings.Join(name[:]," ,"))
}

func main() {
    fmt.Println(strings.Join(os.Args[1:], " "))
    test1()
}



