package main

import "fmt"

//func sum(args ...int)(n int){
func sum(args ...int) int {
	n := 0
	for i := 0; i < len(args); i++ {
		n = n + args[i]
	}
	//return
	return n
}
func main() {
	fmt.Println(sum(1, 2, 3, 4))
}
