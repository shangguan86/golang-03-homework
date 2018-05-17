package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "one two three  one"

	words := strings.Fields(s)

	mcount := make(map[string]int)

	for _, v := range words {
		if _, ok := mcount[v]; ok {
			mcount[v] += 1

		} else {
			mcount[v] = 1
		}

	}

	for k, v := range mcount {

		fmt.Printf("%v : %v\n", k, v)
	}

}
