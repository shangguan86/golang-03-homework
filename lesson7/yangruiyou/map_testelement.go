package main

import "fmt"

func main() {
	var value int
	var isPresent bool

	map1 := make(map[string]int)
	map1["new delhi"] = 55
	map1["bj"] = 20
	map1["sz"] = 30

	value, isPresent = map1["sz"]

	fmt.Println(value, isPresent)

	if isPresent {
		fmt.Printf("the value of sz in map1 is:%d\n", value)
	} else {
		fmt.Printf("map1 does not contain sz.\n")
	}
}
