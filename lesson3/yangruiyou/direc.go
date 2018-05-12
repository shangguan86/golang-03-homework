package main

import (
	"os"
	"log"
	"fmt"
)

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	infos, err := f.Readdir(-2)
	if err != nil {
		log.Fatal(err)
	}

	for _, info := range infos {
		flag := "-"
		if info.IsDir() {
			flag = "d"
		}
		fmt.Printf("%v %dB %s\n", flag, info.Size(), info.Name())

	}
	f.Close()

}
