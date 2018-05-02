package main

import (
	"fmt"
	"flag"
)

func main() {
	//flag.String 返回的是字符串指针
	var help = flag.String("h", "./cat.go", "the file name")
	flag.Parse()
	fmt.Println("The file name is:", *help)
}
