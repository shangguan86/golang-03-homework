package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("请输入3个参数进行计算，你输入的参数是: ", os.Args[1:])
		return
	}

	s1 := os.Args[1]
	op := os.Args[2]
	s2 := os.Args[3]

	n1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("仅支持整数运算，请检查输入，you input is ", "a:", s1, " b:", s2)
		log.Fatal(err)
	}

	n2, err := strconv.Atoi(s2)
	if err != nil {
		fmt.Println("仅支持整数运算，请检查输入，you input is ", "a:", s1, " b:", s2)
		log.Fatal(err)
	}

	switch op {
	case "+":
		fmt.Println("a+b=", n1+n2)
	case "-":
		fmt.Println("a-b=", n1-n2)
	case "*":
		fmt.Println("a*b=", n1*n2)
	case "/":
		if n2 == 0 {
			fmt.Println("Error!,输入的除数是0")
			return
		}
		fmt.Printf("a/b=%.3G\n", float64(n1)/float64(n2))
	default:
		fmt.Println("nothing")
	}

}
