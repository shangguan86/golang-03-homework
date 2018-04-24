package main

import (
	"flag"
	"fmt"
)

func main() {
	//flag.String返回的是字符串指针
	//var file = flag.String("f", "./pointer.go", "the file name!")
	//flag.Parse()                            //使用flag解析参数
	//fmt.Println("the file name is:", *file) //取file指针的值
	//var p string
	//flag.StringVar(&p, "f", "./var.go ", "the file name!")
	//flag.Parse()
	//fmt.Println("flag.StringVar use:", p)

	var m bool
	flag.BoolVar(&m, "n", true, "test bool")
	flag.Parse()
	if m {
		fmt.Println("真")
	} else {
		fmt.Println("假")
	}
}
