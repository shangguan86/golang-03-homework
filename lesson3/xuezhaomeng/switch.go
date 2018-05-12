package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	a := os.Args[1]
	n, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)

	}

	switch n {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3,4,5,6:
		fmt.Println("3")
	default:
		fmt.Println("defa")

	}
}
