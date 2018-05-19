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
	Id   int
	Name string
}

const filename = "./tmp/data.txt"

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
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	con := []uint8(content)
	return con
}

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
			students[name] = Student{Id: id, Name: name}

		case "list":
			for _, v := range students {
				fmt.Printf("Id:%d,Name:%s\n", v.Id, v.Name)
			}

		case "save":
			buf, _ := json.MarshalIndent(students, "", "	")
			fmt.Printf("%s\n", buf)
			writeFile(string(buf))

		case "load":
			buf := readfile(filename)
			//students := make(map[string]Student)
			json.Unmarshal(buf, &students)
			for _,v := range students{
				fmt.Printf("Id:%d,Name:%s\n", v.Id, v.Name)
			}

		case "exit":
			os.Exit(-1)

		case "help":
			continue
		}
	}

}
