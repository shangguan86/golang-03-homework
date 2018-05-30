package main

import (
	"fmt"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func CreateNode(value int) *TreeNode {
	//返回局部变量的地址

	return &TreeNode{Value: value}
}

func (node TreeNode) PrintTree() {
	fmt.Println(node.Value)
}

func (node *TreeNode) SetValue(value int) {
	node.Value = value
}

func main() {
	var root TreeNode
	root = TreeNode{Value: 3}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{5, nil, nil}
	root.Right.Left = new(TreeNode)
	root.Left.Right = CreateNode(2)

	root.Right.Left.SetValue(4)
	root.Right.Left.PrintTree()

	fmt.Println()
}
