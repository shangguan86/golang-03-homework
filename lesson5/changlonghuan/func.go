package main

import (
	"fmt"
)
func add(x,y int)(z int){
	z = x + y
	//裸返回
	return
}
func add1(x int ,y int) int{
	return x+y
}
func add2(x,y int){
	fmt.Println(x+y)
}
func main(){
	fmt.Println(add(22,33))
	fmt.Println(add1(22,33))
	add2(22,33)
}
