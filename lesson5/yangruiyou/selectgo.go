package main

import "fmt"

//冒泡排序
func BubbleSort(b []int) {
	for i := 0; i < len(b); i++ {
		for j := 1; j < len(b)-1; j++ {
			if b[j] < b[j-1] {
				b[j], b[j-1] = b[j-1], b[j]
			}
		}
	}
}

//选择排序

func SelectSort(s []int) {

	for i := 0; i < len(s); i++ {

	}

}

//插入排序
func InsertSort(s []int) {

	for i := 1; i < len(s); i++ {
		for j := 1; j > 0; j-- {
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}

}

//快速排序

func QuickSort(s []int, left, right int) {

	if left >= right {
		return
	}

	val := s[left]
	k := left
	for i := left + 1; i <= right; i++ {
		if s[i] < val {
			s[k] = s[i]
			s[i] = s[k+1]
			k++
		}
	}
	s[k] = val
	QuickSort(s, left, k-1)
	QuickSort(s, k+1, right)

}

func main() {
	bubble := []int{8, 7, 5, 4, 3, 10, 15}
	BubbleSort(bubble)
	fmt.Println(bubble)
}
