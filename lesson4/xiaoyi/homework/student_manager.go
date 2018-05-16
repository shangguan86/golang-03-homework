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

func add(id int, name string) {
	if _, ok := students[name]; ok {
		fmt.Println("Name is exits")
	} else {
		students[name] = Student{Id: id, Name: name}
	}
}

func list() {
	for _, v := range students {
		fmt.Println(v.Id, v.Name)
	}
}

func save() string {
	file, _ := os.Create("/tmp/stu_manager.txt")
	buf01, err := json.Marshal(students)
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(string(buf01))
	file.Sync()
	defer file.Close()

	return "保存成功！"
}

func main() {
	var id int
	var cmd string
	var name string

	//NewScanner创建并返回一个从r读取数据的Scanner，默认的分割函数是ScanLines(行);
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")

		scanner.Scan()
		line := scanner.Text() //Text 将最后一次扫描出的"匹配部分"作为字符串返回(返回副本);
		fmt.Sscan(line, &cmd)

		switch cmd {
		case "add":
			fmt.Sscan(line, &cmd, &id, &name)
			add(id, name)
		case "list":
			list()
		case "save":
			save()
		case "reload":
			f, err03 := ioutil.ReadFile("/tmp/stu_manager.txt")
			if err03 != nil {
				log.Fatal(err03)
			}
			json.Unmarshal(f, &students)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Please enter add/list/save/reload!")

		}
	}
}
