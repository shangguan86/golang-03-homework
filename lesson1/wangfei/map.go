package main

import "fmt"

func maptest()  {
	s := map[string]string {
		"name":"wangfei","age":"26","sex":"female",
	}
	fmt.Println(s)

	m2 := make(map[string]int)
	var m3 map[string]int

	fmt.Println(m2,m3)

	for k,v := range s{
		fmt.Println(k,v)
	}

	fmt.Println("geting values")
	// key 为空的时候 会拿到一个空值 因为 zero values.
	name := s["name"]
	fmt.Println(name)
 }

func main() {
	maptest()
}

