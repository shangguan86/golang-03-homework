package main

import (
	"fmt"
	"flag"
)

func main(){
	//字符串指针
	var file = flag.String("f","./pointer.go","the file name")
	flag.Parse()
	fmt.Println("the file name ",*file)
}