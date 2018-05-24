package main

import "fmt"

func add(m,n int) (r int) {
	r = m + n
	return
}

func main() {
	fmt.Println(add(10,20))
}