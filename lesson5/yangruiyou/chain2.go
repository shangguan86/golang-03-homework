package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

func main() {

	node_a := Node{Val: 1}
	node_b := Node{Val: 2}
	node_c := Node{Val: 3}
	node_d := Node{Val: 4}

	node_a.Next = &node_b

	node_a.Next.Next = &node_c
	node_c.Next = &node_d

	p := &node_a

	printList(p)

	pp := ReverseList(&node_a)
	printList(pp)

}

func ReverseList(p *Node) *Node {
	var pre, next *Node = nil, p.Next
	for p != nil {

		//暂存p的下一个节点
		next = p.Next

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
