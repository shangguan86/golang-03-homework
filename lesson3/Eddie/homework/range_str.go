package main

import (
	"fmt"
)

func main() {
	s := `汉语ab`
	fmt.Println("len of s:", len(s))
	for index, runeValue := range s {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}
