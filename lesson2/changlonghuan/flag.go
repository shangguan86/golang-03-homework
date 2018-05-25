package main

import (
	"flag"
	"fmt"
)

func main() {
	//字符串指针
	var file = flag.String("f", "./cat", "the file name")

	var newlineMark bool
	flag.BoolVar(&newlineMark, "n", false, "是否换行：")

	flag.Parse()

	fmt.Println("the file name ", *file)

	if newlineMark {
		fmt.Println("有换行")
	} else {
		fmt.Println("没有换行")
	}
}
