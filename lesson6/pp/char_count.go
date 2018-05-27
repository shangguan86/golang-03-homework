package main

import (
	//"fmt"
	"io"
	"os"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	b := new(ByteCounter)
	io.Copy(os.Stdout, os.Stdin)
	fmt.Println(*b)
}
