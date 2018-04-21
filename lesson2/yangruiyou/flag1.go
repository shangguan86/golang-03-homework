package main

import (
	"fmt"
	"flag"
)

func main() {
	//flag string
	// var file = flag.String("f", "gopointer.go", "less filename")
	var p string
	//&p 引用传递
	flag.StringVar(&p, "f", "gopointer.go", "less filename")

	var newLineMark bool
	flag.BoolVar(&newLineMark, "n", false, "newLine")

	flag.Parse()
	fmt.Println("filename is:", p)

	fmt.Printf("new line mark: %v\n", newLineMark)

	if newLineMark {
		fmt.Printf("%v\n", newLineMark)
	}

}
