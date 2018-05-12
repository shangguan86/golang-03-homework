package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	//data := []byte("123456789")
	data := os.Args[1]
	f, _ := ioutil.ReadFile(data)

	md5sum1 := md5.Sum(f)
	fmt.Printf("md5:%v,len:%d\n", md5sum1, len(md5sum1))

	filename := os.Args[1]
	name, _ := ioutil.ReadFile(filename)
	md5sum2 := md5.Sum(name)
	fmt.Printf("md5:%v,len:%d", md5sum2, len(md5sum2))
}
