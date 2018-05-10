package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func files(a string) {
	filerf, err := ioutil.ReadFile(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(filerf))
}
func main() {
	var arg = flag.String("f", "", "文件名")
	flag.Parse()
	files(*arg)
}
