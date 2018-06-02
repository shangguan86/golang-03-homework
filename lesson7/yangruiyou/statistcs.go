package main

import "fmt"

func worker(s []int, ch chan int) {
	var result int
	for _, v := range s {
		result += v

	}

	ch <- result
}

func main() {
	var l = []int{1, 3, 6, 4, 0, -9}
	var ch = make(chan int)
	go worker(l[:len(l)/2], ch)
	go worker(l[len(l)/2:], ch)
	x, y := <-ch, <-ch
	fmt.Println(x, y)

}
