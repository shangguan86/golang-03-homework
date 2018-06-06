package main

import (
	"fmt"
	"time"
)

func main() {

	names := []string{"eric", "harry", "robert", "jim", "mark"}
	for _, name := range names {
		go func() {
			fmt.Printf("hello,%s\n", name)
		}()
		time.Sleep(time.Millisecond)

	}
}
