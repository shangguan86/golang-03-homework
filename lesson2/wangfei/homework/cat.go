package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func usage1() {
	fmt.Fprintf(os.Stderr, `[-h help] [-f filename]  Options:`)
	flag.PrintDefaults()
}

func printFile(filename string) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf)
}

// TODO: flag:Parse() 读取命令行参数 -h help, -f filename 接收-f 代表文件名
func main() {

	var h = flag.Bool("h", true, "查看帮助")
	if *h {
		usage1()
	}

	var filename = flag.String("f", "", "指定文件名字")
	flag.Parse()
	if *filename != "" {
		printFile(*filename)
	}

}
