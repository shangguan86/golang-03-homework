package main

import (
	"fmt"
	"unsafe"
)

var (
	v1 int
	v2 int32
	v3 int64
	v4 uint
	v5 uint32
	v6 uint64
)

func main() {

	fmt.Printf("%d\t%d\n", v1, unsafe.Sizeof(v1))
	fmt.Printf("%d\t%d\n", v2, unsafe.Sizeof(v2))
	fmt.Printf("%d\t%d\n", v3, unsafe.Sizeof(v3))
	fmt.Printf("%d\t%d\n", v4, unsafe.Sizeof(v4))
	fmt.Printf("%d\t%d\n", v5, unsafe.Sizeof(v5))
	fmt.Printf("%d\t%d\n", v6, unsafe.Sizeof(v6))

}
