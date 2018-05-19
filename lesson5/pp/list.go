package main

import "fmt"

type Student struct {
    Id int
    Name string
}

type Node struct  {
    Val Student
    Next *Node
}

func main() {
    node_a := Node{Val: Student{Id: 1, Name:"A"}}
    node_b := Node{Val: Student{Id: 2, Name:"B"}}
    //node_c := Node{Val: Student{Id: 3, Name:"C"}}
    //node_d := Node{Val: Student{Id: 4, Name:"D"}}

    node_a.Next = &node_b
    //fmt.Println(node_a)
    p := &node_a
    for p != nil {
        node := *p
        fmt.Println(node)
        p = node.Next
    }
    
}
