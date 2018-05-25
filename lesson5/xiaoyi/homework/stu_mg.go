package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"bufio"
)

type Student struct {
	Id   int
	Name string
}

var students = make(map[string]Student)

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
			reload()
		case "update":
			fmt.Sscan(line,&cmd,&id,&name)
			update(name,id)
		case "del":
			fmt.Sscan(line,&cmd,&name)
			del(name)
		case "exit":
			os.Exit(0)
		default:
			help()

		}
	}
}

// 添加学生信息；
func add(id int, name string) {
	if _, ok := students[name]; ok {
		fmt.Printf("Name %s is exits",name)
	} else {
		students[name] = Student{Id: id, Name: name}
		fmt.Println("add success")
	}
}

// 查看学生信息；
func list() {
	for _, v := range students {
		fmt.Println(v.Id, v.Name)
	}
}

// 保存信息到文件；
func save() {

	file, _ := os.Create("/tmp/stu_manager.txt")

	defer file.Close()

	buf01, err := json.Marshal(students)
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(string(buf01))
	file.Sync()

	fmt.Printf("save information success")
}

// 加载文件里保存的信息；
func reload() {
	f, err := ioutil.ReadFile("/tmp/stu_manager.txt")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(f, &students)
}

// 更改学生信息；
func update(name string,id int) {
	if _,ok := students[name];ok {
		students[name] = Student{id,name}
	} else {
		fmt.Printf("Name %s is don't exits",name)
	}
}

// 删除某个学生的信息;
func del(name string) {
	if _,ok := students[name];ok {
		delete(students,name)
		fmt.Printf("Name %s is delete",name)
	} else {
		fmt.Printf("Name %s don't exits",name)
	}
}

// 查看帮助;
func help() {
	fmt.Println(`
	usage:
	add,添加学生信息,举例：add 1 a;
	del,删除学生信息,举例：del a；
	list,查看所有学生信息;
	save,保存添加学生信息;
	reload,加载某个文件学生信息;
	update,更新某个学生信息,举例：update 4 a;
	exit,退出程序;
`)
}

