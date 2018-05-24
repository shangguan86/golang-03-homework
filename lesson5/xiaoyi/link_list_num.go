package main

import "fmt"

/*
type Student struct {
	Id int
	Name string
}
*/
type Node struct {
	Val int
	Next *Node
}

func printList(p *Node) {
	for p != nil {
		node := *p
		fmt.Println(node)
		p = node.Next
	}
}

func reverseList(p *Node) *Node {
	var pre *Node
	for p != nil {
		// 暂存P的下一个节点
		next := p.Next
		// 让当前节点p的下一个节点为pre
		p.Next =pre
		// 更新pre为当前节点
		pre = p
		// 继续迭代，变更p为原来的next节点
		p = next
	}
	return pre
}

func main() {
	node_a := Node{Val:1}
	node_b := Node{Val:2}
	node_c := Node{Val:3}
	node_d := Node{Val:4}


	node_a.Next = &node_b
	node_a.Next.Next = &node_c
	node_b.Next.Next = &node_d

	//	p := ReverseList(&node_a)

	p := &node_a
	printList(p)

	printList(reverseList(p))
}