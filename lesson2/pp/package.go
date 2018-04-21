package main

import (
	"fmt"
	"os"
)

func print() {
	fmt.Println("abc")
}

func main() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	go print()
}
