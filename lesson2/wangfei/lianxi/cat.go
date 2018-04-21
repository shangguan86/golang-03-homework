package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"flag"
)

func printFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf)
}

func cat_all() {
	filename := os.Args[1]
	fmt.Println(filename)
	//printFile(filename)
}

// TODO: flag:Parse() 读取命令行参数 -h help, -f filename 接收-f 代表文件名
func main() {
	cat_all()
	res := flag.Args()
	fmt.Println(res)
	var ip = flag.Int("flagname", 1234, "help message for flagname")
	fmt.Println(ip)
}
