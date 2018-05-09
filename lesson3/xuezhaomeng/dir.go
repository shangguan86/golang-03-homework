package main

import (
	"fmt"
	"log"
	"os"
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
		flag := "d"
		if ! info.IsDir() {
			flag = "-"
		}
		fmt.Printf("%v  %v  %v %v \n", flag, info.IsDir(), info.Size(), info.Name())

	}

}
