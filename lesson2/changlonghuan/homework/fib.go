package main

import (
	"flag"
	"fmt"
)

//递归方式:运行的过程中调用自己。
func fib(n int) int {
	if n <= 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

//迭代方式：每一次对过程的重复，每一次的结果会作为下一次的初始值。
func fibloop(n int) int {
	//判断n小于2的情况,如果小于2则为n
	if n < 2 {
		return n
	}
	a := 1
	b := 2
	//如果大于2,则为前两次之和
	for i := 2; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
func main() {
	var num1, num2 int
	flag.IntVar(&num1, "n", 10, "输入number1")
	flag.IntVar(&num2, "m", 10, "输入number2")
	flag.Parse()

	fmt.Printf("递归方式:\n  %d的fib结果为：%d\n", num1, fib(num1))
	fmt.Printf("迭代方式:\n  %d的fib结果为：%d\n", num2, fibloop(num2))
}
