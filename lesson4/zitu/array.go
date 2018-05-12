package main

import (
	"log"
	"unsafe"
)

func main() {
	// init1
	var a1 = [3]int{1, 2, 3}
	log.Printf("a1: %v, len: %d, size: %v\n", a1, len(a1), unsafe.Sizeof(a1))
	// init2
	var a2 = [3]int{1, 2}
	log.Printf("a2: %v, len: %d, size: %v\n", a2, len(a2), unsafe.Sizeof(a2))
	// init3
	var a3 = [...]int{1, 2, 3}
	log.Printf("a3: %v, len: %d, size: %v\n", a3, len(a3), unsafe.Sizeof(a3))
	// init4
	var a4 = [...]int{0: 1, 2: 3}
	log.Printf("a4: %v, len: %d, size: %v\n", a4, len(a4), unsafe.Sizeof(a4))

	var b1 [3]int
	b1 = a1
	log.Println(b1 == a1)
	// both length and type should match
	// b2 := [2]int{1, 2}
	// log.Println(b2 == a1)

	// 内存连续
	for i := 0; i < 3; i++ {
		log.Printf("a1[%d]: %d, &a1[%d]: %p, &a1: %p\n", i, a1[i], i, &a1[i], &a1)
	}

	// 2 dim array
	a := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	// a := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	log.Println(a)

	a[1][3] = 1
	row := len(a)
	col := len(a[0])

	log.Println(a, row, col)
}
