package main

import (
    "fmt"
    "strings"
)

func test()  {
    a := [...]int{1,2,3,4,5}
    fmt.Println(a)

    name := []string{"a", "b", "c"}
    fmt.Println(strings.Join(name, ", "))
}

func main() {
    // multiple variable assignment
    //    s, sep := "", ""
    //    fmt.Println(os.Args[0])
    //    for _, arg := range os.Args[1:] {
    //        s += sep + arg
    //        sep = " "
    //    }
    //    fmt.Println(s)
    //}
    test()
}