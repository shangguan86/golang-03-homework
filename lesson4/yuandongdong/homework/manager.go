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

func add(id int, name string) {
	students[name] = Student{Id: id, Name: name}
	fmt.Println("add student ", id, name, "success!")
}

func list() {
	if len(students) == 0 {
		return
	}
	fmt.Println("ID NAME")
	for _, v := range students {
		fmt.Println(v.Id, v.Name)
	}
}

func save(fname string) {
	//	_, err := os.Stat(fname)
	//if os.IsNotExist(err) {
	f, _ := os.Create(fname)
	//} else {
	//	f, _ := os.OpenFile(fname, os.O_APPEND, 0666)
	//}
	defer f.Close()
	data, _ := json.Marshal(students)
	fmt.Fprintf(f, string(data))

}

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
	fmt.Println(err)
	fmt.Println(students)

}
