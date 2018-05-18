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
	var res []byte
	file, _ := os.Open(fname)
	r := bufio.NewReader(file)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		md5sum := md5.Sum([]byte(line))
		for _, v := range md5sum {
			res = append(res, v)
		}
	}
	fmt.Printf("%x\n", res)

}

//对每行做md5sum后的hash值拼接 ！= 读取所有文件做md5sum的hash
