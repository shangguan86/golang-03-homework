package main

import "fmt"

func main() {
	var flist []func()

	for i := 0; i < 3; i++ {
		flist = append(flist, func() {
			a := i
			fmt.Println(a)
		})
	}

	for _, f := range flist {
		f()
	}
}
