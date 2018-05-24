package main

import "fmt"

func add(x, y int) (r int) {
	r = x + y
	// 裸返回
	return
}

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(swap(42, 13))
}
