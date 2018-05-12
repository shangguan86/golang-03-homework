package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func printFile(name string) string {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	return string(buf)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage:\n\t-f string\n\tSource file name after loading the file content (default ' ')\n")
		return
	}

	var s string
	flag.StringVar(&s, "f", " ", "Source file name after loading the file content")
	flag.Parse()

	content := printFile(s)
	fmt.Println(content)
}
