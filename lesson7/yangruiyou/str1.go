package main

import "fmt"

type gopher struct {
	name string
	age  int
}

func main() {
	s := "hello大家好"
	fmt.Println(len(s))
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		fmt.Println(string(r[i]))
	}

	s1 := "hello大家好"
	arr1 := []byte(s1)

	fmt.Println(string(arr1))

	var g = gopher{}
	g.name, g.age = "smile", 20
	fmt.Println(g)

	s2 := make([]int, 0)
	s2 = append(s2, 1, 2, 3)
	fmt.Println(s2)

}
