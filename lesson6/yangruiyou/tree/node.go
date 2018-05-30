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

//使用指针改变结构内容，nil指针可以调用方法

func (node *TreeNode) SetValue(value int) {
	if node == nil {
		fmt.Println("nil node")
	}
	node.Value = value
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.PrintTree()
	node.Right.Traverse()
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

	root.Traverse()
}


//要改变内容必须使用指针接受者
//结构过大考虑使用指针接收者
//值接收者 为go语言特有

//值/指针接收者 均可接收值/指针
