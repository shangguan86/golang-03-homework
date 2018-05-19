package main

import (
	"fmt"
	"strings"
)

//count word function
func CountWord(word []string) map[string]int {
	CountMap := make(map[string]int)
	for _, v := range word {
		if _, ok := CountMap[v]; ok {
			CountMap[v] += 1
		} else {
			CountMap[v] = 1
		}
	}
	return CountMap

}

func main() {
	str := "one two three four one"
	var word []string
	word = strings.Fields(str)

	fmt.Println(CountWord(word))
}
