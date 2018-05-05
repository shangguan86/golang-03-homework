package  main

import (
	"flag"
	"fmt"
)

//func printFile(name string) {
//	buf, err := ioutil.ReadFile(name)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(buf)
//}


func main()  {
	//flag.String 返回的是一个字符串的指针
	//var file = flag.String("f","./pointer.go","the file name")
	//flag.Parse()
	//fmt.Println("the file name is :",*file)
	//filename := *file
	//fmt.Println(filename)

	var newlinemark = flag.Bool("n",false,"print \n")
	flag.Parse()
	fmt.Println(*newlinemark)
	if *newlinemark {
		fmt.Println("\n")
	}

	//var newlinemark bool
	//flag.BoolVar(&newlinemark,"n",false,"show \n")
	//flag.Parse()
	//if newlinemark {
	//	fmt.Println("\n")
	//}

	//var p string
	//flag.StringVar(&p,"h","","show help")
	//flag.StringVar(&p,"f","pointer.go","the file name")
	//// name 选项,value 默认值, uagse 使用说明
	//flag.Parse()
	//fmt.Println("the file name is :",p)
	//filename := p
	//fmt.Println(filename)

	}