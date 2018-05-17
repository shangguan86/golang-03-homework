package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Student struct {
	Id   int
	Name string
}

var students = make(map[string]Student)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(">")
		scanner.Scan()
		line := scanner.Text()
		var cmd string
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "add":
			// add 1 xuexue
			var id int
			var name string
			fmt.Sscan(line, &cmd, &id, &name)
			if name == "" || id == 0 {
				fmt.Println("缺少参数,使用方法: add  Id  Name")
			} else {
				if _, ok := students[name]; ok {
					fmt.Printf("Name:%v 已经存在，请检查后在次添加\n", name)
				} else {
					students[name] = Student{Id: id, Name: name}
				}
			}

		case "list":
			// list
			for _, v := range students {
				fmt.Println(v.Id, v.Name)

			}
		case "save":
			var filename string
			fmt.Sscan(line, &cmd, &filename)
			//buf, err := json.Marshal(students)
			//if err != nil {
			//	log.Fatal("marshal error :%v", err)
			//}
			//f, err := os.Create(filename)
			//defer f.Close()
			//if err != nil {
			//	log.Fatal("Create filename fail")
			//}
			//f.WriteString(string(buf))
			save(students, filename)

		case "load":
			var filename string
			fmt.Sscan(line, &cmd, &filename)
			load(students, filename)

		case "help":
			fmt.Println("--add	:add  id  name")
			fmt.Println("--list	:list")
			fmt.Println("--save	:save filename")
			fmt.Println("--load	:load  filename")
			fmt.Println("--exit	:exit")
		case "exit":
			os.Exit(1)
		default:
			fmt.Println("暂无该命令，请执行help查看帮助")

		}

	}

}



func load(students map[string]Student, filename string) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Read filename fail")
	}
	json.Unmarshal(buf, &students)
	//if err != nil {
	//      log.Fatal("Unmarshl errror:%v", err)
	//}
}

func save(students map[string]Student, filename string) {
	buf, err := json.Marshal(students)
	if err != nil {
		log.Fatal("marshal error :%v", err)
	}
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		log.Fatal("Create filename fail")
	}
	f.WriteString(string(buf))
}
