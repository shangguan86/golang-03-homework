package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return int(*b), nil
}

func main() {
	b := new(ByteCounter)
	io.Copy(b, os.Stdin)
	fmt.Println(*b)
}
