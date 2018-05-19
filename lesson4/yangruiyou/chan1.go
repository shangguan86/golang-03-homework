package main

import "fmt"

func main() {
	counterA := createCouter(2)
	counterB := createCouter(102)
	for i := 0; i < 5; i++ {
		a := <-counterA
		fmt.Printf("(A->%d,B->%d)", a, <-counterB)
	}

}

func createCouter(start int) chan int {
	next := make(chan int)
	go func(i int) {
		for {
			next <- i
			i++
		}
	}(start)
	fmt.Println(next)
	return next
}
