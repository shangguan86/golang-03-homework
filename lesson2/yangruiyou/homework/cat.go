package main

import (
	"io/ioutil"
	"fmt"
	"flag"
)

func printFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Println(string(buf))
}

var filename string

func init() {
	// var help = flag.String("help message", "", "help message")

	flag.StringVar(&filename, "f", "/tmp/a", "print file content")
}

//flag.Parse() 读取命令行参数，接受一个-f的参数，代表文件名，代用printFile()函数打印文件

func main() {

	flag.Parse()
	printFile(filename)

}
