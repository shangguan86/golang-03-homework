package main


type Node struct {
	num  int
	Next *Node
}

func main() {
	node_a := Node{num: 1}
	node_b := Node{num: 2}
	node_c := Node{num: 3}

	node_a.Next = &node_b
	node_a.Next.Next = &node_c

	//P:= ReverseList(&node_a)
	//printList(P)

	p := &node_a

	printList(p)

	printList(reverList(p))
}



func reverList(p *Node) *Node {

	var pre *Node 

	for p != nil {

		next := p.Next
		p.Next = pre
		pre = p
		p = next
	}
	return pre

}
