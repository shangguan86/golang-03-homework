package main

import (
	"fmt"
	"math/rand"
)

//链表尾部插入

type Student struct {
	Name  string
	Age   int
	Score float32

	next *Student
}

func main() {

	var head *Student = new(Student)
	head.Name = "kkgo"
	head.Age = 20
	head.Score = 100

	//var stu1 *Student = &Student{"stu1", 23, 89, nil,}
	//
	//var stu2 *Student = &Student{"stu2", 24, 99, nil,}
	//
	//var stu3 *Student = &Student{"stu3", 25, 79, nil,}
	//
	////fmt.Println(head)
	////fmt.Println(*stu1)
	//
	//head.next = stu1
	//stu1.next = stu2
	//stu2.next = stu3
	//InsertTailChain(&head)
	InsertHeadChain(&head)
	SelectChain(head)

}

func SelectChain(p *Student) {
	for p != nil {
		fmt.Println(*p)
		p = p.next

	}

}

//链表尾部插入
func InsertTailChain(p *Student) {
	var tail = p
	for i := 0; i < 10; i++ {

		stu := &Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		tail.next = stu
		tail = stu
	}
}

//链表头部插入
func InsertHeadChain(p **Student) {
	for i := 0; i < 10; i++ {

		stu := &Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		stu.next = *p
		*p = stu

	}
}

func deleteNodeChain(p *Student) {
	var prev *Student = p
	for p != nil {
		if (p.Name == "stu6") {
			prev.next = p.next
			break
		}
		prev = p
		p = p.next
	}
}

