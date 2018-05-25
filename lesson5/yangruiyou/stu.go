package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student
}

//每个节点包含下一个节点的地址，把链表中的第一个节点叫做链表头

func main() {

	var head *Student
	head.Age = 20
	head.Name = "kkgo"
	head.Score = 80

	insertTail(head)
	trans(head)

}

//遍历单链表

func trans(p *Student) {

	for p != nil {
		fmt.Println(*p)
		p = p.next
	}

}

//尾部插入

func insertTail(p *Student) {

	var tail = p
	for i := 0; i < 10; i++ {
		var stu = Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 10,
		}
		tail.next = &stu
		tail = &stu
	}

}

//单链表头部插入

func insertHead(p *Student) {
	for i := 0; i < 10; i++ {
		var stu = Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 10,
		}
		stu.next = head
	}

}
