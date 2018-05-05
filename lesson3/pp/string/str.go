package main

import (
    "fmt"
)


func main(){
    s1 := "hello" + 
        "world!"

    s2 := `
        hello 
        hello 
        hello 
        hello 
    `
    
    fmt.Println(s1)
    fmt.Println(s2)
}
