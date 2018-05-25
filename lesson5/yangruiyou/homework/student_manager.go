package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"log"
)

type Student struct {
	Id   int
	Name string
}

var stuMap = make(map[string]Student)

func main() {
	run()

}

func run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var cmd string
		var name string
		var id int
		fmt.Print(">")
		scanner.Scan()
		line := scanner.Text()
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "add":
			fmt.Sscan(line, &cmd, &id, &name)
			add(id, name)
		case "del":
			fmt.Sscan(line, &cmd, &name)
			del(name)
		case "update":
			fmt.Sscan(line, &cmd, &id, &name)
			update(id, name)
		case "list":
			list()
		case "save":
			var filename string
			fmt.Sscan(line, &cmd, &filename)
			save(filename)
		case "load":
			var filename string
			fmt.Sscan(line, &cmd, &filename)
			load(filename)
		case "exit":
			return
		default:
			usage()
		}
	}
}

func add(id int, name string) {
	if _, ok := stuMap[name]; ok {
		fmt.Printf("Duplicate name: %s\n", name)
		return
	}
	student := Student{Id: id, Name: name}
	stuMap[name] = student
}

func list() {
	fmt.Printf("ID\tNAME\n")
	for _, stu := range stuMap {
		fmt.Printf("%d\t%s\n", stu.Id, stu.Name)
	}
}

func del(name string) {
	delete(stuMap, name)
}

func save(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	buf, err := json.Marshal(stuMap)
	if err != nil {
		log.Fatal(err)
		return
	}
	f.Write(buf)
}

func load(filename string) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.Unmarshal(buf, &stuMap)
}

func usage() {
	fmt.Println("usage:add|list|del|update|save|load|exit")

}

func update(id int, name string) {
	if _, ok := stuMap[name]; ok {
		del(name)
		add(id, name)
		//stuMap[name] = Student{Id: id, Name: name}
		fmt.Printf("Id :%d update %s\n", id, name)

	} else {
		fmt.Println("user is not exist.")
		add(id, name)
		fmt.Printf("Add Id: %d,Name: %s success.\n", id, name)
	}

}
