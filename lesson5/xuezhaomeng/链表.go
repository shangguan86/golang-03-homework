package main

import "fmt"

type Student struct {
	Id   int
	Name string
	Age  int
	Next *Student
}

func main() {

	head := Student{Id: 1, Name: "xuezm", Age: 19}

	var stu1 Student
	stu1.Id = 2
	stu1.Name = "xuexue"
	stu1.Age = 27

	//返回的是指针
	stu2 := &Student{Id:3, Name:"xuexue", Age:28}


	head.Next = &stu1
	stu1.Next = stu2


	var p *Student = &head
	for p != nil {
		fmt.Println(*p)
		p = p.Next

	}

}
