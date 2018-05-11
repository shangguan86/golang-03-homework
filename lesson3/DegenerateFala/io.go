package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"io"
	"reflect"
)

func main () {
	filename := os.Args[1]
	f, err := os.Open(filename)
	fmt.Println("type of f: ", reflect.TypeOf(f))
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		fmt.Print(line)
		if err == io.EOF {
			break
		}
	}
	f.Close()
}

