package main

import (
	"strings"
	"fmt"
)

func main() {
	s := "one two three one"
	var words []string
	word := make(map[string]int)
	words = strings.Fields(s)

	for _,k := range words {
		/*
		if k == k {
			work[k] += 1
		 */
		if _,ok := word[k];ok {
			word[k] += 1
		}  else {
			word[k] = 1
		}

	}

	for w,c := range word {
		fmt.Printf("%s:%d\n",w,c)
	}
}