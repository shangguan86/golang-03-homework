package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

func main() {
	call()
	//recoverall()
}

func call() {

	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			log.Fatal("False")
		}
	}()

	defer func() { fmt.Println("1") }()
	defer func() { fmt.Println("2") }()
	defer func() { fmt.Println("3") }()
	panic("panic error")

}
