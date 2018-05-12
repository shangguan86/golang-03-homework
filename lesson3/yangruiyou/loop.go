package main

import "fmt"

func toUpper(s string) string {
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 'a' && arr[i] <= 'z' {
			arr[i] = arr[i] - 32
		}
	}
	return string(arr)
}

func main() {
	n := 1
	if n > 2 {
		fmt.Println("big")

	} else if n < 2 && n != 1 {
		fmt.Println("small")

	} else {
		fmt.Println("equal")
	}
	s1 := "hello中国"
	fmt.Print(toUpper(s1))
}
