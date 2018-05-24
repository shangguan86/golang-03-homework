package main

import (
	"fmt"
)

type Student struct {
	Id   int
	Name string
}

type Node struct {
	Val  Student
	Next *Node
}

func main() {

	node_a := Node{Val: Student{Id: 1, Name: "student_a"}}
	node_b := Node{Val: Student{Id: 2, Name: "student_b"}}
	node_c := Node{Val: Student{Id: 3, Name: "student_c"}}
	node_d := Node{Val: Student{Id: 4, Name: "student_d"}}

	node_a.Next = &node_b

	node_a.Next.Next = &node_c
	node_c.Next = &node_d

	p := &node_a

	printList(p)

	pp := ReverseList(&node_a)
	printList(pp)

}

func ReverseList(p *Node) *Node {
	//var pre, next *Node = nil, p.Next

	var pre *Node = nil
	for p != nil {

		//暂存p的下一个节点
		next := p.Next

		//让当前节点p的下一个节点为pre

		p.Next = pre
		//更新pre为当前点
		pre = p
		//继续迭代，变更p为其原来的next节点
		p = next
	}
	//此时，p是nil,应当返回pre节点
	return pre
}

func printList(p *Node) {

	for p != nil {
		node := *p
		fmt.Println(node)
		p = node.Next
	}

}
