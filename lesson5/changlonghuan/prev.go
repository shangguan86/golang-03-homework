package main

func reverseList(p *Node) *Node{
	var prev *Node
	for p != nil {
		//暂存p的下一个节点
		next := p.Next
		//让当前节点p的下一个节点为prev
		p.Next = prev
		//更新prev为当前节点
		prev = p
		//继续迭代,变更p为其原来的next节点
		p = next
	}
	return prev
}
