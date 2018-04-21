package main

import (
	"fmt"
	"flag"
)

func main() {
	//flag string
	var file = flag.String("f", "gopointer.go", "less filename")

	flag.Parse()
	fmt.Println("filename is:", *file)
}
