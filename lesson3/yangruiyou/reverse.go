package main

import "fmt"

func Reverse(s string) string {
	a := []rune(s)
	for k, v := range a {
		fmt.Print(k, v)
	}

	s2 := string(a)
	return s2
}
