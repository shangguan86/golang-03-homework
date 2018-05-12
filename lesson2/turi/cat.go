package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func printFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))

}

func main() {

}
