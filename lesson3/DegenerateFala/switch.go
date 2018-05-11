package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
)

func main () {
	s := os.Args[1]

	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	switch n {
	case 1:
		fmt.Println("a")
	case 2:
		fmt.Println("a")
	case 3:
		fmt.Println("b")
	case 4:
		fmt.Println("d")
	default:
		fmt.Println("noting")
	}
}
