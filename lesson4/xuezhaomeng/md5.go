package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	op := os.Args[1]

	buf, err := ioutil.ReadFile(op)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(buf))

	data := []byte(string(buf))
	md5sum := md5.Sum(data)
	fmt.Printf("md5:%x, len:%d\n", md5sum, len(md5sum))
	// %x ,  [size]byte,  const size=16  
}
