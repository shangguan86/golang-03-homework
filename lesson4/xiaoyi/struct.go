package main

import "fmt"

type student struct {
	Id int
	Name string
}

func main() {
	var s student
	// struct 使用方式1
	s.Id = 1
	s.Name = "tom"
	fmt.Println(s)

	// struct 使用方式2
	s1 := student{
		Id:2,
		Name:"alice",
	}
	fmt.Println(s1)
	// struct 指针,通过该种方式可以改变struct的值;
	var p *student // 方式1
	p = &s1
	p.Id =3
	fmt.Println(s1)

	var p1 *int // 方式2
	p1 = &s1.Id
	*p1 = 4
	fmt.Println(s1)
}