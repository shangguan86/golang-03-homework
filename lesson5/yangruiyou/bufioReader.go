package main

import (
	"os"
	"bufio"
	"io"
)

func main() {
	var f *os.File
	f = os.Stdin
	r := bufio.NewReader(f)
	for {
		line, e := r.ReadString('\n')
		if e == io.EOF {
			break
		}
		os.Stdout.Write([]byte(line))
	}
}
