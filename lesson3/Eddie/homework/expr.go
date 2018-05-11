package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func strToint(s string) int {
	si, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return si
}
func strTofla(s string) float64 {
	si, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return si
}

func main() {
	s1 := os.Args[1]
	op := os.Args[2]
	s2 := os.Args[3]

	switch op {
	case "+":
		n1 := strToint(s1)
		n2 := strToint(s2)
		fmt.Printf("%d+%d=%d\n", n1, n2, n1+n2)
	case "-":
		n1 := strToint(s1)
		n2 := strToint(s2)
		fmt.Printf("%d-%d=%d\n", n1, n2, n1-n2)
	case "*":
		n1 := strToint(s1)
		n2 := strToint(s2)
		fmt.Printf("%d*%d=%d\n", n1, n2, n1*n2)
	case "/":
		n1 := strTofla(s1)
		n2 := strTofla(s2)
		if n2 != 0 {
			fmt.Printf("%v / %v = %v  \n", n1, n2, n1/n2)
		} else {
			fmt.Println("除数不可为0!\n")
		}
	default:
		fmt.Println("nothing")
	}
}
