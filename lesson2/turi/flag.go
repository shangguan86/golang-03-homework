package main

import (
	"flag"
	"fmt"
)

func main() {
	var file = flag.String("f", "./pointer.go", "the file name")

	var newlineMark bool
	flag.BoolVar(&newlineMark, "n", false, "newline")
	flag.Parse()
	fmt.Println("the file name is", *file)
	if newlineMark {
		fmt.Println("newline")
	}
}
