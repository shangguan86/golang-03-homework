package main

import "fmt"

//定义结构体
type Student struct {
	Id   int
	Name string
}

func main() {
	// 1.  声明s 为结构体
	var s Student
	// 赋值
	s.Id = 1
	s.Name = "jack"
	fmt.Println(s)



	// 2.  
	s1 := Student{
		Id:   2,
		Name: "alice",
	}
	fmt.Println(s1)

	s1 = s
	fmt.Println(s1)
}
