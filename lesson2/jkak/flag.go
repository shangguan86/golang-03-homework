package main

import (
	"flag"
	"fmt"
)

func main() {
	// return a string pointer
	var file = flag.String("f", "./pointer.go", "description of pointer file")

	var p string
	flag.StringVar(&p, "k", "./var.go", "description of var file")
	// 将命令行参数-f后面的值复制到p变量里面

	var n = flag.Bool("n", false, "enter new line")

	flag.Parse()

	fmt.Println("file name:", *file)
	fmt.Println("var file:", p)

	if *n {
		fmt.Println("\tnew line")
	}
}
