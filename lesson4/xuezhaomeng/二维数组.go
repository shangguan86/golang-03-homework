package main

import "fmt"

func main() {
	a := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	a[1][3] = 1
	row := len(a)
	clo := len(a[0])
	fmt.Println(a, row, clo)
	fmt.Println(a[0],a[1])

}
