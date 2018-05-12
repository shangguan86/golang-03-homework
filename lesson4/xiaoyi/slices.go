package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	/*
		slice 为什么从 0 开始取值？
		p -> a
		p +  0
		p +  1*sozepf(type)
		p +  2*sozeof(type)
	*/

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for i := range a {
		fmt.Printf("%d %d\n", i, a[i])
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var r [3]int = [3]int{1, 2, 3}
	fmt.Printf("r:%d,len:%d,size:%v\n", r, len(r), unsafe.Sizeof(r))

	var d [3]int = [3]int{1, 2}
	fmt.Printf("a:%d,len:%d,size:%v\n", d, len(d), unsafe.Sizeof(d))

	b := [...]int{1, 3, 4}
	b[2] = 0
	fmt.Printf("b:%d,len:%d,size:%v\n", b, len(b), unsafe.Sizeof(b))

	c := [...]int{0: 1, 2: 3}
	fmt.Printf("c:%d,len:%d,size:%v\n", c, len(c), unsafe.Sizeof(c))

}
