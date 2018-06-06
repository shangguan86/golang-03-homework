package main

import (
	"time"
)

func main() {

	go println("go!goroutine!")
	time.Sleep(time.Millisecond)
}
