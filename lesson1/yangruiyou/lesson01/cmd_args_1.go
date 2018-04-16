package main

import (
	"os"
	"fmt"
)

func main() {
	var s, sep string
	s="-"

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "-"
	}
	fmt.Println(s)
}
