package main

import "fmt"

type Student struct {
	Id   int
	Name string
}

type Node struct {
	Val  Student
	Next *Node
}

func reverseList(p *Node) *Node {
	var pre *Node
	for p != nil {
		// 暂存 p 的下一个节点
		next := p.Next
		// 让当前节点p的下一个节点为pre
		p.Next = pre
		// 更新 pre为当前节点
		pre = p
		// 继续迭代， 变更 p 为其原来的next节点
		p = next
	}
	// 此时，p是nil, 应当返回pre节点
	return pre
}

func main() {
	fmt.Println(x)

	node_a := Node{Val: Student{Id: 1, Name: "A"}}
	node_b := Node{Val: Student{Id: 2, Name: "B"}}
	//node_c := Node{Val: Student{Id: 3, Name:"C"}}
	//node_d := Node{Val: Student{Id: 4, Name:"D"}}

	node_a.Next = &node_b
	//fmt.Println(node_a)
	p := &node_a
	printList((p))
	fmt.Println()
	printList(reverseList(p))
}
