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

func main() {
	var filename string
	flag.StringVar(&filename, "f", "./cat file_to_path", "读取内容")
	flag.Parse()
	printFile(filename)
}
