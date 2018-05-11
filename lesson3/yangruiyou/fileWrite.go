package main

import (
	"os"
	"log"
)

func main() {

	//f, err := os.Create("f.txt")
	f, err := os.OpenFile("f.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("hello world.\n")
	f.Close()
}
