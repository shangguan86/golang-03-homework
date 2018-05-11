package main

import (
	"fmt"
	"os"
	"strconv"
)

var Usage = func() {
	fmt.Println("Usage: go run expr.go <interger1> [+,-,*,/] <interger2>")

	fmt.Println("\nThe commands are:\n\t+\tAddition of two values.\n\t-\tSubtract of two values\n\t*\tMultiply of two values\n\t/\tDivide of two values")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 4 {
		Usage()
		return
	}
	switch args[2] {
	case "+":
		v1, err1 := strconv.Atoi(args[1])

		v2, err2 := strconv.Atoi(args[3])

		if err1 != nil || err2 != nil {
			fmt.Println("Usage:go run expr.go <interger1> + <interger2>")
			return
		}

		fmt.Println("Result:", v1+v2)

	case "-":
		v1, err1 := strconv.Atoi(args[1])

		v2, err2 := strconv.Atoi(args[3])

		if err1 != nil || err2 != nil {
			fmt.Println("Usage:go run expr.go <interger1> - <interger2>")
			return
		}

		fmt.Println("Result:", v1-v2)

	case "*":
		v1, err1 := strconv.Atoi(args[1])

		v2, err2 := strconv.Atoi(args[3])

		if err1 != nil || err2 != nil {
			fmt.Println("Usage:go run expr.go <interger1> * <interger2>")
			return
		}

		fmt.Println("Result:", v1*v2)

	case "/":
		v1, err1 := strconv.Atoi(args[1])

		v2, err2 := strconv.Atoi(args[3])

		if err1 != nil || err2 != nil {
			fmt.Println("Usage:go run expr.go <interger1> / <interger2>")
			return
		}

		if v2 == 0 {
			fmt.Println("error: 0 cannot be the dividend!")
			return
		}

		fmt.Println("Result:", v1/v2)

	default:
		Usage()
	}
}
