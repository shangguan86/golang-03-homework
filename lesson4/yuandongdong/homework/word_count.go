package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "one two three one"
	var words []string
	myMap := make(map[string]int)

	words = strings.Fields(s)

	for _, v := range words {
		if _, ok := myMap[v]; ok {
			myMap[v] += 1
		} else {
			myMap[v] = 1
		}

	}

	for k, v := range myMap {
		fmt.Printf("\"%s\" count is %d\n", k, v)
	}

}
