package main

import (
	"flag"
	"fmt"
)

func main() {
	var flagNewLine bool
	filePath := flag.String("f", "default path", "file path")
	flag.BoolVar(&flagNewLine, "n", false, "new line flag")
	flag.Parse()
	fmt.Println(flagNewLine)

	// printFile(*filePath)
	fmt.Println(*filePath)
	if flagNewLine {
		fmt.Println()
	}
}
