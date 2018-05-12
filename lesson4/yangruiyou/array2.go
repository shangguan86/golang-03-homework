package main

import (
	"fmt"
	"unsafe"
	"unicode/utf8"
)

func main() {
	str := "计算字符长度kkgo"

	fmt.Println(len(str), utf8.RuneCountInString(str), len([]rune(str)))

	bb := "我爱美丽的中国kkgo"
	fmt.Println(utf8.RuneCountInString(bb), len([]rune(bb)), len(bb))

	aa := [4]int{1: 2, 3: 4}
	for k, v := range aa {
		fmt.Println(k, v)
	}
	var r [3]int = [3]int{1, 2, 3}
	fmt.Printf("r:%v,len:%d,size:%v\n", r, len(r), unsafe.Sizeof(r))

	var a [3]int = [3]int{1, 2}

	fmt.Printf("a:%v,len:%d,size:%v\n", a, len(a), unsafe.Sizeof(a))

	b := [...]int{1, 2, 3}
	b[2] = 0
	fmt.Printf("b:%v,len:%d,size:%v\n", b, len(b), unsafe.Sizeof(b))

}
