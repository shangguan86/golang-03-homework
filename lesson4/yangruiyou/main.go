package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
)

func main() {
	if len(os.Args) < 2 {
		return

	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return

	}
	defer file.Close()

	reader := bufio.NewReader(file)

	//ReadLine(line,isPrefix,err)
	var line int

	for {
		_, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}

		if !isPrefix {
			line++
		}

	}
	fmt.Println(line)
}
