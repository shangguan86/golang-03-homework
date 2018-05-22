package main

import "fmt"

type Student struct {
	Id   int
	Name string
	Age  int
	Next *Student
}

func trans(p *Student) {

	for p != nil {
		fmt.Println(*p)
		p = p.Next

	}

}

func InsertTail(p *Student) {
	var tail = p
	for i := 0; i < 10; i++ {

		stu := Student{Id: i, Name: "xue", Age: i}

		tail.Next = &stu
		tail = &stu

	}

}

//func InsertHead(p *Student) {
//	var head = p
//
//	for i := 0; i < 10; i++ {
//
//		stu := Student{Id: i, Name: "xuexye", Age: i}
//
//		stu.Next = head
//		head = &stu
//	}
//
//}

func main() {

	//使用指针的方式定义
	var head *Student = &Student{}
	//var head *Student = new(Student)

	//链尾插入
	//InsertTail(head)
	//trans(head)

	fmt.Println()

	//链头插入
	for i := 0; i < 10; i++ {
		stu := Student{Id: i, Name: "xue", Age: i}
		stu.Next = head
		head = &stu
	}
	trans(head)
}
