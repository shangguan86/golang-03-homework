package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {

	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++

	}
	//fmt.Println(keys)
	sort.Strings(keys)
	//fmt.Println(keys)

	for _, value := range keys {

		fmt.Printf("  key: %v value:%d / ", value, barVal[value])

	}
}
