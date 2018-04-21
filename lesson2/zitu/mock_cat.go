package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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
	filePath := flag.String("f", "", "file path")
	flag.Parse()
	if *filePath == "" {
		log.Fatal("need -f flag")
	}
	printFile(*filePath)
}
