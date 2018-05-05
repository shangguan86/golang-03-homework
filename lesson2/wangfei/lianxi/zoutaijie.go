package main

import (
	"fmt"
	"flag"
	"time"
)

func fib1(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return fib1(n-1) + fib1(n-2)
}

func fib2(n int) int {
	var (
		x     int
		y     int
		count int
	)

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	x = 1
	y = 2
	for i := 2; i < n; i++ {
		count = x + y
		x, y = y, count
	}
	return count
}

func main() {
	var num int
	flag.IntVar(&num, "n", 0, "the num of fib")
	flag.Parse()

	start1 := time.Now()
	var res = fib1(num)
	t1 := time.Now()
	elapsed1 := t1.Sub(start1)
	fmt.Printf("递归算法,num:%d,计算结果:%d,耗时:%v\n", num,res,elapsed1)

	start2 := time.Now()
	var res2 = fib2(num)
	t2 := time.Now()
	elapsed2 := t2.Sub(start2)
	fmt.Printf("循环算法,num:%d,计算结果:%d,耗时:%v\n",num, res2,elapsed2)
}
