// package
package main

// std lib
import (
    "fmt"
    "math/rand"
    "time"
)

// the entrance


func generate_random() int {

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

    rand_num := rand.Intn(100)
    //fmt.Printf("随机数是:%d\n",rand_num)
    return  rand_num
}


func main() { // go fmt
    fmt.Println("hello golang")
    // support Unicode
    fmt.Println("hello go语言")
    res := generate_random()
    fmt.Println(res)
}
