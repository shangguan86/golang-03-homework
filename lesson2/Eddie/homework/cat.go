package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

// START OMIT
func printFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}

// END OMIT
func main() {
	var fileName string
	flag.StringVar(&fileName, "f", "default fileName", "help msg for fileName")
	flag.Parse()
	fmt.Println(len(flag.Args()))
	for i, v := range flag.Args() {
		fmt.Println(i, v)
		printFile(v)
	}
}
