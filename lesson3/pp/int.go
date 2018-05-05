package main

import (
    "fmt"
    "unsafe"
    "reflect"
)

func main() {
    var (
        v1 int
        //v2 int32
        //v3 int64
        //v4 uint
        //v5 uint32
        //v6 uint64
    )

    fmt.Printf("v1 is %d, size: %d\n", v1, unsafe.Sizeof(v1))
    fmt.Println(reflect.TypeOf(v1))
}
