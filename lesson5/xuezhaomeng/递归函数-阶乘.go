package main

import "fmt"

func j(n int) int {
	if n > 0 {
		result := n * j(n-1)
		return result
	}
	return 1
}

func main() {

	fmt.Println(j(15))

}
