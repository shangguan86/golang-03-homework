package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func printFile(name string) {
	fmt.Println(name)
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))

}

func main() {
	//f := os.Args[1]
	var p string
	flag.StringVar(&p, "f", "", "Input File Name")
	flag.Parse()
	//fmt.Println(p)
	printFile(p)

}
