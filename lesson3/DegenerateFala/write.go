package main

import (
	"os"
)

func main() {
	//f,_ := os.Create("a.txt")
	f, _ := os.OpenFile("a.txt", os.O_CREATE | os.O_APPEND | os.O_RDWR, 0644)
	f.WriteString("helloworld!\n")
	f.Seek(2, os.SEEK_SET)
	f.WriteString("$$$")
	f.Close()
	}
