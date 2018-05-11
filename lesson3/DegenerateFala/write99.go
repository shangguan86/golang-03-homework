package main

import (
	"os"
	"fmt"
)

func main () {
	f, _ := os.Create("table.txt")
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d  ", i, j, i*j)
			fmt.Fprintf(f, "%d * %d = %d ", i, j, i*j)
		}
		fmt.Println()
		fmt.Fprintln(f)
	}
	f.Close()
}
