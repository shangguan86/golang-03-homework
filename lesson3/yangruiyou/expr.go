package main

import (
	"os"
	"strconv"
	"log"
	"fmt"
)

func main() {
	s1 := os.Args[1]
	op := os.Args[2]
	s2 := os.Args[3]

	n1, err := strconv.Atoi(s1)

	if err != nil {
		log.Fatal(err)
	}

	n2, err := strconv.Atoi(s2)
	if err != nil {
		log.Fatal(err)
	}

	switch op {
	case "+":
		fmt.Println("a+b", n1+n2)
	default:
		fmt.Println("Nothing")

	}

}
