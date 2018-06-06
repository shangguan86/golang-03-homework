package main

import (
	"fmt"
	"unsafe"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

type struct2 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 20.5
	ms.str = "kkgo"

	fmt.Printf("%v,%v,%v\n", ms.i1, ms.f1, ms.str)

	fmt.Println(ms)

	m2 := &struct2{2, 4.5, "yang"}
	fmt.Println(*m2)

	size := unsafe.Sizeof(m2)
	fmt.Println(size)
}
