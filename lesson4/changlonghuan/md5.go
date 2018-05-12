package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := ioutil.ReadFile(os.Args[1])
	data := []byte(file)
	md5sum := md5.Sum(data)

	fmt.Printf("%x,%d\n", md5sum, len(md5sum))
}
