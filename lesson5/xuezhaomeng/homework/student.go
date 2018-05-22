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
			add(id, name, students)
		case "list":
			// list
			list(students)
		case "save":
			//save filename
			var filename string
			fmt.Sscan(line, &cmd, &filename)
			save(students, filename)

		case "load":
			//load filename
			var filename string
			fmt.Sscan(line, &cmd, &filename)
			load(students, filename)
		case "delete":
			//delete name
			var name string
			fmt.Sscan(line, &cmd, &name)
			del(students, name)
		case "update":
			var name string
			var newid int
			fmt.Sscan(line, &cmd, &name, &newid)
			update(students, name, newid)

		case "help":
			help()
		case "exit":
			os.Exit(1)
		default:
			fmt.Println("暂无该命令，请执行help查看帮助")

		}

	}

}
func list(students map[string]Student) {

	for _, v := range students {
		fmt.Println(v.Id, v.Name)

	}

}

func add(id int, name string, students map[string]Student) {

	if name == "" || id == 0 {
		fmt.Println("缺少参数,使用方法: add  Id  Name")
	} else {
		if _, ok := students[name]; ok {
			fmt.Printf("Name:%v 已经存在，请检查后在次添加\n", name)
		} else {
			students[name] = Student{Id: id, Name: name}
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

func del(students map[string]Student, name string) {
	delete(students, name)

}

func update(students map[string]Student, name string, newid int) {
	if _, ok := students[name]; ok {
		fmt.Printf("key 存在，将要改变ID为新值 %v\n", newid)

	} else {

		fmt.Printf("key 不存在，新增Id:%v,Name:%v\n ", newid, name)

	}
	students[name] = Student{Id: newid, Name: name}

}

func help() {

	fmt.Println("--add	:add  id  name")
	fmt.Println("--list	:list")
	fmt.Println("--save	:save filename")
	fmt.Println("--load	:load  filename")
	fmt.Println("--exit	:exit")
	fmt.Println("--delete	:delete name")
	fmt.Println("--update	:update name newid")

}
