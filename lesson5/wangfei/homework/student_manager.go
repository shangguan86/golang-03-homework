package main

import (
	"os"
	"log"
	"fmt"
	"io/ioutil"
	"bufio"
	"encoding/json"
)

type Student struct {
	Id   int    `json:"student_id"`
	Name string `json:"student_name"`
}

const filename = "/tmp/data.txt"

var students = make(map[string]Student)

func writeFile(s string) {
	file1, err := os.Create(filename)
	defer file1.Close()
	if err != nil {
		fmt.Println(file1, err)
		return
	}
	file1.WriteString(s)
}

func readfile(filename string) []uint8 {
	//一次性读取到内存 适合小文件读取
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	con := []uint8(content)
	return con
}

func usage() {
	fmt.Printf(`
Usage:
  add:   添加信息
  list:  列出所有用户信息
  save:  保存用户信息到/tmp/data.txt
  update: 更新信息
  delete: 删除信息
  exit:  退出系统
  help:  查看帮助
`)
}

func add(id int, name string) {
	students[name] = Student{Id: id, Name: name}
}

func list() {
	for _, v := range students {
		fmt.Printf("Id:%d,Name:%s\n", v.Id, v.Name)
	}
}

func save() {
	buf, _ := json.MarshalIndent(students, "", "	")
	fmt.Printf("%s\n", buf)
	writeFile(string(buf))
}

func loads() {
	buf := readfile(filename)
	json.Unmarshal(buf, &students)
	for _, v := range students {
		fmt.Printf("Id:%d,Name:%s\n", v.Id, v.Name)
	}
}

func remove(name string) {
	delete(students, name)
}

func update(id int, name string) {

	if len(name) == 0 {
		fmt.Printf("name is null\n")
		return
	}

	if _, ok := students[namfe]; ok {
		add(id, name)
		fmt.Printf("update %s is successed!\n", name)
		list()
	} else {
		fmt.Printf("%s is not exist!\n", name)
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()
		var cmd string
		fmt.Sscan(line, &cmd)
		var id int
		var name string

		switch cmd {

		case "help":
			usage()

		case "add":
			fmt.Sscan(line, &cmd, &id, &name)
			add(id, name)

		case "list":
			list()

		case "save":
			save()

		case "delete":
			fmt.Sscan(line, &cmd, &name)
			remove(name)

		case "update":
			fmt.Sscan(line, &cmd, &id, &name)
			update(id, name)

		case "load":
			loads()
		case "exit":
			os.Exit(-1)

		}
	}

}
