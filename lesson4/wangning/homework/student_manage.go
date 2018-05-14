package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Student struct {
	Id   int
	Name string
}

var students = make(map[string]Student)
var dbfile = "students.db"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		scanner.Scan()
		line := scanner.Text()
		var cmd string
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "add":
			// add 1 Tom
			var id string
			var uid int
			var name string
			fmt.Sscan(line, &cmd, &id, &name)
			uid, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Id not a number")
				continue
			}
			add(uid, name)
		case "list":
			// list:
			// ID NAME
			// 1  TOM
			// 2  JACK
			list()
		case "load":
			// read from file
			// json.UnMarshal -> students
			status, err := load(dbfile)
			if !status {
				fmt.Println(err)
				continue
			}
			fmt.Println("Load File success ", dbfile)
		case "save":
			// json.Marshal  students -> string
			// write to file
			r_status, err := save(dbfile, students)
			if !r_status {
				fmt.Println(err)
				continue
			}
			fmt.Println("Save success ", dbfile)
		case "help":
			// print help message
			Help()
		case "flush":
			// print help message
			flush()
		case "exit":
			os.Exit(1)
		default:
			// do some stuff..
			Help()
		}
	}
}

func save(dbfile string, stu map[string]Student) (status bool, err error) {
	buf, err := json.Marshal(stu)
	if err != nil {
		return false, err
	}
	//fmt.Println(string(buf))
	var w1 = []byte(string(buf))
	err2 := ioutil.WriteFile(dbfile, w1, 0666)
	if err2 != nil {

		return false, err2

	}
	return true, nil
}

func load(dbname string) (status bool, err error) {

	f, err := ioutil.ReadFile(dbname)
	if err != nil {
		return false, err
	}
	var s map[string]Student
	s = make(map[string]Student)
	err1 := json.Unmarshal([]byte(f), &s)
	if err1 != nil {
		return false, err1
	}
	students = s
	return true, nil
}

func add(id int, name string) {
	if _, ok := students[name]; ok {
		fmt.Printf("Student %s  has already existed!\n", name)
	} else {
		students[name] = Student{Id: id, Name: name}
	}
}

func list() {
	for _, v := range students {
		fmt.Printf("%d\t%s\n", v.Id, v.Name)
	}
}
func flush() {
	s := make(map[string]Student)
	students = s
	fmt.Printf("学生系统内存数据已被清理,%s\n",students)

}

func Help() {
	fmt.Println("参数说明如下:")
	flag.Usage()

}

var (
	h bool

	adduser, showlist, loadfile, savefile, flushdata bool
)

func init() {
	flag.BoolVar(&adduser, "add", false, "Add Student Notice,Usage:  add 1 XXX")
	flag.BoolVar(&savefile, "save", false, "Save Students Edit,Usage:  save")
	flag.BoolVar(&showlist, "list", false, "Show Students List,Usage:  list")
	flag.BoolVar(&loadfile, "load", false, "Reload Students Database,Usage: load")
	flag.BoolVar(&flushdata, "flush", false, "临时清理所有用户数据,需save保存,Usage: flush")
	flag.BoolVar(&h, "help", false, "View help information,Usage: help")
	flag.Usage = usage

}

func usage() {
	flag.PrintDefaults()
}
