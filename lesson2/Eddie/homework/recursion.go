package main

import (
	"flag"
	"fmt"
	"time"
)

//递归计算
func fib(n int) int {
	var res int
	if n < 2 {
		res = n
	} else {
		res = fib(n-1) + fib(n-2)
	}
	return res
}

//计算裴波那契数列
func main() {
	var num int
	flag.IntVar(&num, "n", 0, "Please input a number!")
	flag.Parse()
	args := flag.NArg()
	if len(flag.Args()) < 0 {
		fmt.Println(args)
		fmt.Println("Please input a number!")
	}
	for i := 0; i <= num; i++ {
		startTime := time.Now()
		var stepsNum int = fib(i)
		endTime := time.Now()
		useTime := endTime.Sub(startTime)
		fmt.Printf("台阶数为%d: 可以有 %d 种走法。递归计算耗时为:%s ,共计算:%d 次\n", i, stepsNum, useTime, i)
	}
}
