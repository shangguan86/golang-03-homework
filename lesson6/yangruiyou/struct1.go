package main

import "fmt"

//结构字段内存布局是连续的

type Student struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	var stu Student
	stu.Name = "kkgo"
	stu.Age = 18
	stu.Score = 90

	fmt.Println(stu)

	fmt.Printf("name:%p\n", &stu.Name)
	fmt.Printf("age：%p\n", &stu.Age)
	fmt.Printf("score:%p\n", &stu.Score)

	//结构体初始化

	var stu1 *Student = &Student{
		Name:  "kk",
		Age:   20,
		Score: 99,
	}
	fmt.Println(*stu1)

	var stu3 = Student{
		Name: "kk3",
		Age:  20,
	}
	fmt.Println(stu3)
}
