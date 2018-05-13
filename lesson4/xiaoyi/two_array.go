package main

import "fmt"

func main() {
	a := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

	a[1][3] = 1 //将二维数组索引为1的数组的第3个索引值改为1;
	row := len(a)
	col := len(a[0])
	fmt.Println(a, row, col)
}
