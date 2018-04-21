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
//读flag参数
func Parse() string {
	flag.Parse()
	args := flag.Args()
	return args[0]
}

func main() {
	printFile(Parse())

}
