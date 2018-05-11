package main

import (
	"flag"
	"fmt"
	"time"
)

func fib(n int) int {
	stepNum, stepsNum := 0, 1
	for i := 0; i < n; i++ {
		stepNum, stepsNum = stepsNum, stepNum+stepsNum
	}
	return stepNum
}

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
		fmt.Printf("台阶数为%d: 可以有 %d 种走法。迭代计算耗时为:%s ,共计算:%d 次\n", i, stepsNum, useTime, i)
	}
}
