package main

import (
	"io/ioutil"
	"fmt"
)

func printFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Println(string(buf))
}
