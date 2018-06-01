package main

import (
	"fmt"
	"strings"
	"sort"
)


type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return len(p[i]) < len(p[j]) }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }



func main() {

	str := "hello  go src 51reboot"

	s := strings.Fields(str)

	sort.Sort(StringSlice(s))
	fmt.Println(s)
}


