package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func printFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Println(string(buf))
}

func r_flag() {
	Fn := flag.String("f", "/tmp/aa", "Enter a filename")
	flag.Parse() //注意, 如果想获取命令行的指定值, 必须要写入该行
	fmt.Println(*Fn)
	printFile(*Fn)
}

func main() {
	r_flag()
}
