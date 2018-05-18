package main

import (
	"bufio"
	"os"
	"fmt"
)

type student struct {
	//Id   int `json:"student_id" `
	Id   int
	Name string
}

var students = make(map[string]student)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(">>>>>input")
		scanner.Scan()
		line := scanner.Text()
		var cmd string
		fmt.Scan(line, &cmd)

		switch cmd {

		case "add":
			var id int
			var name1 string
			fmt.Scan(line, &cmd, &name1)
			students[name1] = student{Id: id, Name: name1}
			//调用方法
		case "list":
			fmt.Printf("list command")
			for _, v := range students {
				fmt.Printf("name:%v,id:%v", v.Name, v.Id)
			}
		case "save":
			//wtire to file
			//json.Marshal()
		case "load":
			//read from file
			//json.Unmarshal()
		case "help":
			//
		case "exit":
			fmt.Println("exit")
			break
		}

	}

}
