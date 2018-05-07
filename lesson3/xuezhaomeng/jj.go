package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("table.txt")
	for i := 1; i < 10; i++ {

		for ii := 1; ii <= i; ii++ {

			//fmt.Printf(" %v * %v = %v ", i, ii, i*ii)
			fmt.Fprintf(f, " %v * %v = %v ", i, ii, i*ii)

		}
		//fmt.Println()
		fmt.Fprintf(f, "\n")
	}

}

