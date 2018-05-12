package main

import (
	"flag"
	"fmt"
	"time"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	var num int
	flag.IntVar(&num, "n", 0, "Please input a number!")
	flag.Parse()
	for i := 0; i <= num; i++ {
		startTime := time.Now()
		var f = fib(i)
		endTime := time.Now()
		useTime := endTime.Sub(startTime)
		fmt.Printf("台阶数是 %d 可以有: %d 种走法. 计算耗时是：%s\n", i, f, useTime)
	}
}
