package main

import (
	"os"
	"crypto/md5"
	"fmt"
)

func main() {
	file := []byte(os.Args[1])

	md5sum := md5.Sum([]byte(file))
	fmt.Printf("md5:%x,len:%d\n",md5sum,len(md5sum))

}
