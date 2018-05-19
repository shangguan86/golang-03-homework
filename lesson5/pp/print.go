package main

import "fmt"

var x = 1

func printList(p *Node) {
	for p != nil {
		node := *p
		fmt.Println(node)
		p = node.Next
	}
}
