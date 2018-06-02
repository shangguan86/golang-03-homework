package main

import (
	"fmt"
	"sort"
)

type StringSlice1 []string

func (p StringSlice1) Len() int           { return len(p) }
func (p StringSlice1) Less(i, j int) bool { return len(p[i]) < len(p[j]) }
func (p StringSlice1) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Sort(StringSlice1(s))
	fmt.Println(s)
}