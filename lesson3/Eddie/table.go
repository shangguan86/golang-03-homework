package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("table.txt")
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Fprintf(f, "%d*%d=%-4d", i, j, i*j)
		}
		fmt.Println()
		fmt.Fprintln(f)
	}
	f.Close()
}
