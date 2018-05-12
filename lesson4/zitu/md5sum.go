package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	path := os.Args[1]
	f, _ := ioutil.ReadFile(path)
	// f, _ := os.Open(path)
	md5sum := md5.Sum(f)
	fmt.Printf("md5:%x, len:%d\n", md5sum, len(md5sum))
}
