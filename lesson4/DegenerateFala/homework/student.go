package main

import (
	"fmt"
	"bufio"
	"os"
)

type Student struct {
	Id int
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
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "add":
			// add 1 Tom
			var id int
			var name string
			fmt.Sscan(line, &cmd, &id, &name)
			add(id, name)
		case "list":
			// list:
			// ID NAME
			// 1  TOM
			// 2  JACK
		case "load":
			// read from file
			// json.UnMarshal -> students
		case  "save":
			// json.Marshal  students -> string
			// write to file
		case "help":
			// print help message
		case "exit":
			// exit
		default:
			// do some stuff..
		}
	}
}

func add(id int, name string) {
	students[name] = Student{Id: id, Name: name}
}