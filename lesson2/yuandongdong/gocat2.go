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

//TODO:flag.Parse()
//通过flag.Args读取filename
func Parse() string {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("no argument or more than one argument")
		return ""
	}
	return args[0]
}

func main() {
	printFile(Parse())

}
