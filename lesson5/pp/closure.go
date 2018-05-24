package main

import "fmt"

func main() {
	var funcs []func()

	for i := 0; i < 3; i++ {
		var x = i
		funcs = append(funcs, func() {
			fmt.Println(x)
		})
	}

	// i : 0
	// i : 1
	// i : 2
	// i == 3

	for _, f := range funcs {
		f()
	}
}
