package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args:= os.Args
	if len(args) < 4 {
		log.Fatal("Args Count < 3")
	}
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
		fmt.Printf("%d + %d = %d\n", n1,n2,n1+n2)
	case "-":
		fmt.Printf("%d = %d = %d\n", n1,n2,n1-n2)
	case "*":
		fmt.Printf("%d * %d = %d\n", n1,n2,n1*n2)
	case "/":
		if n2 == 0 {
			fmt.Println("除数不能等于0")
		} else if n2 > n1 {
			fmt.Printf("%d / %d = %-.2f\n", n1,n2,float64(n1)/float64(n2))
		} else {
			fmt.Printf("%d / %d = %d\n", n1,n2,n1/n2)
		}
	default:
		fmt.Println("nothing")
	}

}
