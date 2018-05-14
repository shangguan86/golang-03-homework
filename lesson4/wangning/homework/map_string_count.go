package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "one two three one"
	var words []string
	var maps map[string]int
	maps = make(map[string]int)
	words = strings.Fields(s)

//	fmt.Println(words)
	for _, v := range words {

		//fmt.Println(v)
		if _, ok := maps[v]; ok {
			maps[v] += 1
		} else {
			maps[v] = 1
		}
	}
	for k,v := range maps {
		fmt.Println(k,v)
	}

}
