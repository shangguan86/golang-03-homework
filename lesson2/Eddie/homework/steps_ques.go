package main

import (
	"flag"
	"fmt"
	"time"
)

//递归计算
func fibre(n int) int {
	var res int
	if n < 2 {
		res = n
	} else {
		res = fibre(n-1) + fibre(n-2)
	}
	return res
}

func fibit(n int) int {
	stepNum, stepsNum := 0, 1
	for i := 0; i < n; i++ {
		stepNum, stepsNum = stepsNum, stepNum+stepsNum
	}
	return stepNum
}

func main() {

	//计算裴波那契数列

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
		var stepsNumre int = fibre(i)
		endTime := time.Now()
		useTime := endTime.Sub(startTime)
		fmt.Printf("台阶数为%d: 可以有 %d 种走法。递归计算耗时为:%s ,共计算:%d 次\n", i, stepsNumre, useTime, i)
		startTime1 := time.Now()
		var stepsNumit int = fibit(i)
		endTime1 := time.Now()
		useTime1 := endTime1.Sub(startTime1)
		fmt.Printf("台阶数为%d: 可以有 %d 种走法。迭代计算耗时为:%s ,共计算:%d 次\n", i, stepsNumit, useTime1, i)
	}
}
