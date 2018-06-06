package main

import (
	"fmt"
	"time"
)

func main() {
	name := "eric"
	go func() {
		fmt.Printf("hello,%s\n", name)
	}()
	time.Sleep(time.Millisecond)

	name = "harry"
	time.Sleep(time.Millisecond)





}
