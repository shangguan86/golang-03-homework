package main

import (
	"strings"
	"fmt"
)

func main() {
	s := "one two three one"
	var words []string
	words = strings.Fields(s)

	word := make(map[string]int)
	for _, v := range words {
		if v == v {
			word[v] += 1
		}
	}

	for w,c := range word {
		fmt.Printf("%s:%d\n",w,c)
	}
}