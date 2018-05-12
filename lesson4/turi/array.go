package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for k, v := range a {
		fmt.Printf("%d %d\n", k, v)
	}

	for i := range a {
		fmt.Printf("%d %d\n", i, a[i])

	}

	for _, v := range a {
		fmt.Printf("%d \n", v)
	}

}
