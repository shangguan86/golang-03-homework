package main

import (
	"os"
	"strconv"
	"log"
	"fmt"
)

func main() {
	s := os.Args[1]
	n, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}

	switch n {
	case 1:
		fmt.Println("a")
	case 2:
		fmt.Println("b")
	case 3:
		fmt.Println("c")
	case 4:
		fmt.Println("d")
	default:
		fmt.Println("nothing")
	}

	switch {
	case n == 1:
		fmt.Println("A")
	case n == 2:
		fmt.Println("B")
	case n == 3:
		fmt.Println("C")
	case n == 4:
		fmt.Println("D")
	default:
		fmt.Println("NOTHING")
	}

}
