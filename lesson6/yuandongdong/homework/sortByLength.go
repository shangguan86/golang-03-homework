package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return len(p[i]) < len(p[j]) }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	s := []string{
		"Ab",
		"ccc",
		"cdasfvr",
		"B",
	}
	fmt.Println("Before sort", s)
	sort.Sort(StringSlice(s))

	fmt.Println("After sort", s)
}
