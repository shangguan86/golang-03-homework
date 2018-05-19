package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("a.txt")
	defer f.Close()
	buf := make([]byte, 1024)
	n, _ := f.Read(buf)
	fmt.Println(string(buf[:n]))
}
