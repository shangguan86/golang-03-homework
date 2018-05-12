package main

import "fmt"

func main() {
    a1 := [3]int{1,2}
    //a2 := [2]int{1,2}
    //fmt.Println(a1 == a2)

    for i:=0;i<3;i++ {
        fmt.Printf("a1[%d]:%d, &a1[%d]: %p, &a1: %p\n", i, a1[i], i, &a1[i], &a1) 
    }

    // p -> a1
    // &a1[0]   p + 0*sizeOf(int)
    // &a1[1]   p + 1*sizeOf(int)
    // &a1[2]   p + 2*sizeOf(int)
}

