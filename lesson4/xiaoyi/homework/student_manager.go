package main

import (
	"bufio"
	"os"
	"fmt"
)

type Student struct {
	Id int
	Name string
}

var students = make(map[string]Student)

func add(id int,name string) {
	students[name] = Student{Id:id,Name:name}
}

func main() {
	var id int
	var cmd string
	var name string

	//NewScanner创建并返回一个从r读取数据的Scanner，默认的分割函数是ScanLines(行);
	scanner := bufio.NewScanner(os.Stdin)
	for  {
		fmt.Print(">")

		scanner.Scan()
		line := scanner.Text() //Text 将最后一次扫描出的"匹配部分"作为字符串返回(返回副本);
		fmt.Sscan(line,&cmd)

		switch cmd {
		case "add":
			fmt.Sscan(line,&cmd,&id,&name)
			add(id,name)
		case "list":
			for _,v := range students {
				fmt.Println(v.Id,v.Name)
			}
		}
	}
}