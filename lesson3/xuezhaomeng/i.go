package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var (
		v1 int
		v2 int32
		v3 int64
		v4 uint
		v5 uint32
		v6 uint64
	)

	fmt.Printf("v1 is %d ,siez %d\n", v1, unsafe.Sizeof(v1))
	fmt.Printf("v2 is %d ,siez %d\n", v2, unsafe.Sizeof(v2))
	fmt.Printf("v3 is %d ,siez %d\n", v3, unsafe.Sizeof(v3))
	fmt.Printf("v4 is %d ,siez %d\n", v4, unsafe.Sizeof(v4))
	fmt.Printf("v5 is %d ,siez %d\n", v5, unsafe.Sizeof(v5))
	fmt.Printf("v6 is %d ,siez %d\n", v6, unsafe.Sizeof(v6))

}
