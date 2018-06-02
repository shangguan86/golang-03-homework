package main

import (
	"fmt"
)

type gopher struct {
	name string
	age  int
}

func main() {

	//var g *gopher = &gopher{name: "zhang", age: 20,}
	//fmt.Println(*g)

	var m = make(map[string]*gopher)

	var l = []gopher{
		{name: "tome", age: 11},
		{name: "ami", age: 12},
		{name: "li", age: 13},
	}
	for i, r := range l {

		【'"""'
		//fmt.Println(l[0])

		if r.name == "tom" {
			l[i].name = "wang"
		}
		m[r.name] = &l[i]

	}

	for k, v := range m {
		fmt.Println(k, v.name, v.age)
	}

}
