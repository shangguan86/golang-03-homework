package main

import "fmt"

func main() {

    
    a := [4]int{1: 2, 3:4}

    for k, v := range a {
        fmt.Println(k, v) 
    }

}
