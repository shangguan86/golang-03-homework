package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	fname := os.Args[1]
	var res string
	file, _ := os.Open(fname)
	r := bufio.NewReader(file)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		res = res + line
	}
	md5sum := md5.Sum([]byte(res))
	fmt.Printf("%x\n", md5sum)

}
