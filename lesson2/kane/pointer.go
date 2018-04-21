package main

import "fmt"

func main() {
	// 打印地址
	a := 10
	fmt.Println(*&a)

	//申明一个纸箱int类型的数据指针
	p := &a
	*p = 20
	fmt.Println(a)
}
