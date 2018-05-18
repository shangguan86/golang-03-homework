package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "one two there one five two four"
	var words []string
	words = strings.Fields(s)
	bufWord := make(map[string]int)
	for _, v := range words {
		bufWord[v] += 1
	}
	//fmt.Println(bufWord)
	fmt.Printf("word\tcount\n")
	for k, v := range bufWord {
		fmt.Printf("%v\t%d\n",k, v)
	}
}
