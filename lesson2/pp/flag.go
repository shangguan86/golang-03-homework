package main

import (
    "fmt"
    "flag"
)

func main() {
    //var file = flag.String("f", "./pointer.go", "the file name")

    var p string
    flag.StringVar(&p, "f", "./pointer.go", "the go file name")

    var newlineMark bool
    flag.BoolVar(&newlineMark, "n", false, "the new line mark")

    flag.Parse()

    //fmt.Println("the file name is: ", *file)
    fmt.Println("the file name is: ", p)

    fmt.Printf("new line mark: %t\n", newlineMark)

    if newlineMark {
        fmt.Println() 
    }
}
