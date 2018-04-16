package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(s)
}
