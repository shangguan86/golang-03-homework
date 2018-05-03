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
	fmt.Println(buf)
	fmt.Printf("%U", 10)
	fmt.Println(string(buf))
}

//TODO:flag.Parse()
//读取-f参数
func Parse() string {
	fname := flag.String("f", "", "file name")
	flag.Parse()
	if fname == nil {
		fmt.Println("no argument")
	}
	return *fname
}

func main() {
	f := Parse()
	printFile(f)

}
