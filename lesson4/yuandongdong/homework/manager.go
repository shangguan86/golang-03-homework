package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()
		var cmd string
		var fname string
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "add":
			// add 1 Tom
			var id int
			var name string
			fmt.Sscan(line, &cmd, &id, &name)
			add(id, name)
		case "list":
			list()
		case "load":
			// read from file
			// json.UnMarshal -> students
			fmt.Sscan(line, &cmd, &fname)
			load(fname)
		case "save":
			// json.Marshal  students -> string
			// write to file
			fmt.Sscan(line, &cmd, &fname)
			save(fname)
		case "help":
			// print help message
			fmt.Println(`
			usage:
			add id name
			list
			save filename
			load filename
			exit
			`)
		case "exit":
			// exit
			os.Exit(1)
		default:
			// do some stuff..
		}
	}
}

//添加学生信息
func add(id int, name string) {
	//处理id为0或name为空的情况
	if id == 0 {
		fmt.Println("invalid id: ", id, "; plz check you input and try again.")
		return
	} else if len(name) == 0 {
		fmt.Println("lack student name, plz check you input and try again.")
		return
	}
	//判断student name是否重复
	if _, ok := students[name]; ok {
		fmt.Println("add student fail; ", name, "already exist.")
		return
	}
	students[name] = Student{Id: id, Name: name}
	fmt.Println("add student ", id, name, "success!")
}

//列出学生信息
func list() {
	if len(students) == 0 {
		fmt.Println("no student infomation")
		return
	}
	fmt.Println("ID NAME")
	for _, v := range students {
		fmt.Println(v.Id, v.Name)
	}
}

//保存学生信息到文件
func save(fname string) {
	var file *os.File
	//没有增加学生信息，不需要保存
	if len(students) == 0 {
		return
	}

	if checkFileIsExist(fname) {
		//防止学生信息被覆盖,save之前先load学生信息
		load(fname)
		file, _ = os.OpenFile(fname, os.O_RDWR, 0666)
	} else {
		file, _ = os.Create(fname)
	}
	defer file.Close()
	data, _ := json.Marshal(students)
	fmt.Fprint(file, string(data))

}

//从文件中读取学生信息
func load(fname string) {
	var data []byte
	file, err := os.Open(fname)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ = ioutil.ReadAll(file)
	err = json.Unmarshal(data, &students)
	fmt.Println("load students infomation success!")

}

//判断filename文件是否存在
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
