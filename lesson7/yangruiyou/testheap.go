package main

import "fmt"

//go tool compile -S  testheap.go

//go tool compile -v  testheap.go

var g := "smile"

func test1() {
	var a [1]int
	c := a[:]
	fmt.Println(c)
}

func test2() {
	var a [1]int
	c := a[:]
	fmt.Println(c)
}
