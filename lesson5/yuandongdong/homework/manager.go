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
	var cmd, name string
	var id int
	var fname string
	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "add":
			// add 1 Tom
			fmt.Sscan(line, &cmd, &id, &name)
			add(id, name)
		case "update":
			fmt.Sscan(line, &cmd, &id, &name)
			update(id, name)
		case "delete":
			fmt.Sscan(line, &cmd, &name)
			del(name)
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
			help()
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

	file, err := os.OpenFile(fname, os.O_RDWR|os.O_TRUNC, 0666)
	checkError(err)
	defer file.Close()
	data, err := json.Marshal(students)
	checkError(err)
	fmt.Fprint(file, string(data))
	fmt.Println("save file", fname, "success!")

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
	data, err = ioutil.ReadAll(file)
	checkError(err)
	err = json.Unmarshal(data, &students)
	checkError(err)
	fmt.Println("load file ", fname, " success!")

}

//判断filename文件是否存在
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//check error
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func help() {
	fmt.Println(`
	usage:
	add id name: 新增学生
	update id name：修改学生ID
	delete name：删除学生
	list：列出所有学生
	save filename：保存学生信息到文件中
	load filename：从文件中加载学生信息
	exit:退出
	`)
}

//更新学生信息
func update(id int, name string) {
	if _, ok := students[name]; ok {
		students[name] = Student{
			id,
			name,
		}
		fmt.Println("update student information success. new student informantion is :", id, name)
	} else {
		fmt.Println("student ", name, "do not exist.")
		return
	}
}

//删除学生信息
func del(name string) {
	if _, ok := students[name]; !ok {
		fmt.Println("student ", name, "do not exist.")
		return
	}
	delete(students, name)
	fmt.Println("delete ", name, "success.")
}
