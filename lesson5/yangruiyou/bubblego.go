package main

import "fmt"

//冒泡排序
func Bubble(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-1; j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

func main() {
	bubble := []int{8, 7, 5, 4, 3, 10, 15}
	Bubble(bubble)
	fmt.Println(bubble)
}
