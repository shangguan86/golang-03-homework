package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("a.txt")

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Fprintf(f, "%d * %d = %2d ", i, j, i*j)
		}
		fmt.Fprintf(f, "\n")
		//		fmt.Fprintln(f)
	}

	f.Close()
}
