package main

import "fmt"

func main() {

	a := []string{
		"one",
		"two",
		"three",
	}

	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
	}

	fmt.Println(a)

}
