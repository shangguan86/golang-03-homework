package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "one two three one"
	var words []string
	var mapscount map[string]int
	mapscount = make(map[string]int)
	words = strings.Fields(s)
	//  fmt.Println(maps)
	//	fmt.Println(words)

	for _, x := range words {

		if _, ok := mapscount[x]; ok {
			mapscount[x] += 1
		} else {
			mapscount[x] = 1
		}
	}
	for y,x := range mapscount {
		fmt.Printf("%v : %v个\n",y,x)
	}

}