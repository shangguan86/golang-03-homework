package main

import (
	"flag"
	"fmt"
)

func main() {
	var p string
	var newLineMark bool
	flag.StringVar(&p, "f", "./pointer.go", "the file name")
	flag.BoolVar(&newLineMark, "n", false, "the bool vale")
	flag.Parse()
	fmt.Println("value of newLineMark:", newLineMark)
	if newLineMark {
		fmt.Println("---")
		fmt.Printf("\n")
		fmt.Println("---")
	}
}
