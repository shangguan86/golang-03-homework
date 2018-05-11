package main

import "fmt"

func main() {

	var b = false

	fmt.Printf("%v\n", b)

	var a = &b
        *a = true
    	
	fmt.Printf("%v\n", b)


}
