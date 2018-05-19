package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("a.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf, _ := ioutil.ReadAll(f)
	fmt.Println(string(buf))
}
