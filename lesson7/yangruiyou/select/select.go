package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {

	var c1, c2 = generator(), generator()

	w:=createWorker(0)
	hasValue:=false


	for {
		n:=0

		select {
		case n := <-c1:
			//fmt.Println("received from c1:", n)
			hasValue=true
		case n := <-c2:
			//fmt.Println("received from c2:", n)
			//default:
			//	fmt.Println("no values")
			hasValue=true

			case w<-n:
				if hasValue{

				}

		}

	}

}

func generator() chan int {

	out := make(chan int)
	go func() {

		for {
			i := 0

			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++

		}

	}()

	return out
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}
