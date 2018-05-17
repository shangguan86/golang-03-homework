package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// 1.
	var a [3]int = [3]int{1, 2, 3}

	fmt.Printf("%v\n", a[1])

	// 2.

	var ab [3]int = [3]int{1, 2}
	fmt.Printf("a: %v, len: %d, size: %v\n", ab, len(ab), unsafe.Sizeof(ab))
	// 3.

	b := [...]int{1, 2, 3,4,5}
	b[2] = 0
	fmt.Printf("b: %v, len: %d, size: %v\n", b, len(b), unsafe.Sizeof(b))

	//4.

	c := [...]int{0: 1, 2: 3}
	fmt.Printf("c: %v, len: %d, size: %v\n", c, len(c), unsafe.Sizeof(c))

}
