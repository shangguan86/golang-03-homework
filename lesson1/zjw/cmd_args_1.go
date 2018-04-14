package main

// import syntax
import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	// for loop(the only loop statement in Go):
	// for initialization; condition; post {}

	// while loop:
	// for condition {}

	// infinite loop：
	// for {}

	for i := 1; i < len(os.Args); i++ {
		// --i is illegal in Go
		sep = "-"
		s += sep + os.Args[i]
	}
	fmt.Println(s)
}
