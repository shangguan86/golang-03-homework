package main

import "fmt"

func printList(p *Node){
	for p != nil{
		node := *p
		fmt.Println(node.Val.Id,p.Next)
		p = node.Next
	}
}
