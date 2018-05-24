package main

import (
	"fmt"
)

type Student struct {
	Id   int
	Name string
}

type Node struct {
	Val  int
	Next *Node // 试试看改成Node
}

func main() {
	node_a := Node{Val: 1}
	node_b := Node{Val: 2}
	node_c := Node{Val: 3}
	node_d := Node{
		Val: 4,
	}

	node_a.Next = &node_b
	node_a.Next.Next = &node_c
	node_c.Next = &node_d

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
