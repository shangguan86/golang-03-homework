package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	s := os.Args[1]

	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	switch n {
	case 1, 2:
		fmt.Println("a")
	case 3:
		fmt.Println("b")
	case 4:
		fmt.Println("d")
	default:
		fmt.Println("nothing")
	}

	fmt.Println()

	switch {
	case n == 1:
		fmt.Println("a")
	default:
		fmt.Println("b")
	}

}
