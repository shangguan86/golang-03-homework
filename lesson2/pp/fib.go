package main

import (
    "fmt"
    "flag" // 获取数列的n
)

func fib(n int) int {
    // n==1 
    if n == 1 {
        return 1 
    }

    // n==2
    if n == 2 {
        return 2 
    }

    // n > 2
    // return fib(n-1) + fib(n-2)
    f1 := fib(n-1)
    f2 := fib(n-2)
    fn := f1 + f2
    return fn
}

func main() {
    var num int
    flag.IntVar(&num, "n", 0, "the number of fibonacci")
    flag.Parse()

    var f = fib(num)
    fmt.Printf("fibonacci of %d is: %d\n", num, f)

}


