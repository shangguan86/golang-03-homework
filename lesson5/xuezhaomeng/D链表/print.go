package main

import "fmt"

func printList(p *Node) {
	for p != nil {

		node := *p
		fmt.Println(node)
		p = node.Next

	}
}
