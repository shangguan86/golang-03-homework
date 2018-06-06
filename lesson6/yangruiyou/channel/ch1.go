package main

import "fmt"

func chanDemo() {
	c := make(chan int)

	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	//想c发送数据
	c <- 1
	c <- 2

	n := <-c
	fmt.Println(n)

}

func mina() {
	chanDemo()
}
