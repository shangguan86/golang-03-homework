package main

import (
	"bufio"
	"fmt"
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
		fmt.Println(">")
		scanner.Scan()
		line := scanner.Text()
		var cmd string
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "add":
			// add 1 xuexue
			var id int
			var name string
			fmt.Sscan(line, &cmd, &id, &name)
			students[name] = Student{Id: id, Name: name}

		case "list":
			// list
			for _,v := range students {
			fmt.Println(v.Id,v.Name)
			
			}
		}

	}

}
