package main

import (
	"flag"
	"fmt"
)

func main() {

	//var file = flag.String("f", "readme.me", "Input something")
	//flag.Parse()
	//fmt.Println(*file)

	//var p string
	var newlineMark bool
	//flag.StringVar(&p, "f", "./pointer.go","input")
	flag.BoolVar(&newlineMark, "n", false, "input")
	flag.Parse()


	fmt.Printf("%t", newlineMark)
	//if newlineMark{
	//fmt.Println()
	//}else{
	//	fmt.Println("not 换行")
	//}

}
