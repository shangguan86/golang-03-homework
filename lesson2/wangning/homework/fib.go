package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)

}

func fib1(n int) int {

	var i, j int
	i, j = 1, 1
	for m := 0; m <= n-2; m++ {

		i, j = j, i+j

	}

	return i

}

func NowTime() time.Time {
	timestamp := time.Now().Unix()
	now_time := time.Unix(timestamp, 0)
	return now_time

}

func main() {
	var num int
	args := os.Args
	if args == nil || len(args) < 2 {
		log.Fatal("Please Input a Args")
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("-----recursion:BEGIN Time:%s-----\n", NowTime())
	fmt.Printf("Return Value is: %d\n", fib(num))
	fmt.Printf("-----recursion:END Time:%s-----\n", NowTime())
	fmt.Println()
	fmt.Println()
	fmt.Printf("-----No recursion:BEGIN Time:%s-----\n", NowTime())
	fmt.Printf("Return Value is: %d\n", fib1(num))
	fmt.Printf("-----No recursion:END Time:%s-----\n", NowTime())
}
