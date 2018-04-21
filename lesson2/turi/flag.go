package main

import (
	"flag"
	"fmt"
)

func main() {
	var file = flag.String("f", "./pointer.go", "the file name")
	flag.Parse()
	fmt.Println("the file name is", *file)

}
