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
	Next *Node // 试试看改成Node
}

func main() {
	node_a := Node{Val: Student{Id: 1, Name: "student_a"}}
	node_b := Node{Val: Student{Id: 2, Name: "student_b"}}
	node_c := Node{Val: Student{Id: 3, Name: "student_c"}}

	node_a.Next = &node_b
	node_a.Next.Next = &node_c

	p := &node_a
	fmt.Println("before reverse:")
	printList(p)
	fmt.Println()
	fmt.Println("after reverse:")
	p = ReverseList(&node_a)
	printList(p)
}

func ReverseList(cur *Node) *Node {
	var prevNode *Node
	for cur != nil {
		nextNode := cur.Next
		cur.Next = prevNode
		prevNode = cur
		cur = nextNode

	}
	return prevNode

}

func printList(p *Node) {
	for p != nil {
		fmt.Printf("Node address: %p, Node value: %v \n", p, *p)
		p = p.Next
	}
}
