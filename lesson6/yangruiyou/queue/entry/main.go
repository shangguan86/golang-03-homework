package main

import (
	"github.com/51reboot/golang-03-homework/lesson6/yangruiyou/queue"
	"fmt"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q)
	fmt.Println(q.Pop())
	q.IsEmpty()
}
