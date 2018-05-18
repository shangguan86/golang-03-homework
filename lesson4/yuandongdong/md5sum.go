package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fname := os.Args[1]

	filedata, _ := ioutil.ReadFile(fname)
	md5sum := md5.Sum(filedata)
	fmt.Printf("%x\n", md5sum)

}
