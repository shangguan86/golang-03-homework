package main

import (
	"fmt"
)

func main() {
	Array_a := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	slice_a := Array_a[2:5]
	//cap()函数返回的是数组切片分配的空间大小
	fmt.Printf("len=%d cap=%d %v\n", len(slice_a), cap(slice_a), slice_a)

	slice_b := slice_a[0:8]
	fmt.Printf("len=%d cp=%d %v\n", len(slice_b), cap(slice_b), slice_b)
}
