package main

import (
	"fmt"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 14}
	var s []int = primes[1:4]
	fmt.Println(s)
	/*
	   slice 为什么从 0 开始取值？
	   p -> a,p 是指针的意思
	   p +  0
	   p +  1*sozepf(type)
	   p +  2*sozeof(type)
	*/
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	// 切片值更改对整个slice有效;
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	//打印a,b共同元素
	for _, k := range a {
		for _, v := range b {
			if k == v {
				fmt.Println(k, v)
			}
		}
	}
	// 切片方法
	x := []int{2, 3, 5, 7, 11, 13}
	x[0] = 1

	x = x[0:4]
	fmt.Println(x)

	x = x[:2]
	fmt.Println(x)

	x = x[1:]
	fmt.Println(x)

	x1 := x
	fmt.Println(x1 == nil)
	//切片长度(len)和容量(cap,它表示的是切片开始到整个slice底层的长度)
	Array_a := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	Slice_a := Array_a[2:5]
	fmt.Printf("len=%d cap=%d %v\n", len(Slice_a), cap(Slice_a), Slice_a)

	Slice_b := Slice_a[0:8]
	fmt.Printf("len=%d cap=%d %v\n", len(Slice_b), cap(Slice_b), Slice_b)
}