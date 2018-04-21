package main

import (
	"fmt"
)

func main() {
	//指向int类型数据的指针
	var a = 10
	var b = true
	var c = "fun"
	var d = 100
	var e = &d

	//声明指针
	var p *int
	//a的内存地址
	p = &a
	//a的内存地址重新赋值
	*p = 1000
	fmt.Printf("%v\n%v\n%v\n%v\n%v",*&a, &b, &c, *e, *p)
}
