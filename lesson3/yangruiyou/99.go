package main

import (
	"fmt"
	"os"
	"log"
)

func main() {

	f, err := os.OpenFile("f.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			// fmt.Printf("%d * %d = %d ", i, j, i*j)
			fmt.Fprintf(f, "%d * %d = %d ", i, j, i*j)
			//f.WriteString("%d * %d = %d ", i, j, i*j)

		}
		// fmt.Println()
		fmt.Fprintln(f)

	}

}
