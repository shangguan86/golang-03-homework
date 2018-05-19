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
		fmt.Sscan(line, &cmd)

		switch cmd {
		case "add":
			var id int
			var name string

			fmt.Sscan(line, &cmd, &id, &name)

			add(id, name)

		case "list":
			list()

		case "remove":
			var name string
			fmt.Sscan(line, &cmd, &name)
			remove(name)

		case "save":
			var file string
			fmt.Sscan(line, &cmd, &file)
			save(file)

		case "load":
			var file string
			fmt.Sscan(line, &cmd, &file)
			load(file)

		case "exit":
			os.Exit(1)

		default:
			Usage()

		}

	}

}

func add(id int, name string) {
	if id < 1 {
		fmt.Printf("%c[%d;%d;%dmStudent id Must be a positive integer that is non-zero!%c[0m\n", 0x1B, 0, 40, 31, 0x1B)
		return
	}

	if name == "" {
		fmt.Printf("%c[%d;%d;%dmStudent name cannot be empty!%c[0m\n", 0x1B, 0, 40, 31, 0x1B)
		return
	}

	if _, ok := students[name]; ok {
		fmt.Printf("%c[%d;%d;%dmStudent is exist!%c[0m\n", 0x1B, 0, 40, 31, 0x1B)
		return
	}

	students[name] = Student{
		Id:   id,
		Name: name,
	}

	fmt.Printf("%c[%d;%d;%dmAdd student %s success!%c[0m\n", 0x1B, 1, 40, 32, name, 0x1B)
	return
}

func list() {
	fmt.Print("+---------------------+\n")
	fmt.Printf("|%-10s|%-10s|\n", "ID", "NAME")
	fmt.Print("+---------------------+\n")

	for _, v := range students {
		fmt.Printf("|%-10d|%-10s|\n", v.Id, v.Name)
	}

	fmt.Print("+---------------------+\n")
}

func remove(name string) {
	if name == "" {
		fmt.Printf("%c[%d;%d;%dmStudent name cannot be empty!%c[0m\n", 0x1B, 0, 40, 31, 0x1B)
		return
	}

	if _, ok := students[name]; !ok {
		fmt.Printf("%c[%d;%d;%dmStudent is not exist!%c[0m\n", 0x1B, 0, 40, 31, 0x1B)
		return
	}

	delete(students, name)

	fmt.Printf("%c[%d;%d;%dmRemove student %s success!%c[0m\n", 0x1B, 1, 40, 32, name, 0x1B)
	return
}

func save(file string) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	d, err := json.Marshal(students)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(f, string(d))

	fmt.Printf("%c[%d;%d;%dmSave student info success!%c[0m\n", 0x1B, 1, 40, 32, 0x1B)
	return

}

func load(file string) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(b, &students); err != nil {
		panic(err)
	}

	fmt.Printf("%c[%d;%d;%dmLoad student info success!%c[0m\n", 0x1B, 1, 40, 32, 0x1B)
	return
}

func Usage() {
	fmt.Println("Usage:\n\tadd <Id:interger> <Name:string>\n\tlist <show students info!>\n\tremove <Name:string>\n\tload <filename>\n\tsave <filename>")
}
