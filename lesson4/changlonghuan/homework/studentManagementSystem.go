package main

import (
	"bufio"
	 "encoding/json"
	"fmt"
	"log"
	"os"
	"io/ioutil"
)

type Student struct {
	Id   int
	Name string
}

var students = make(map[string]Student)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
//func save(filename string) {
//	file, err := os.Create("StudentSystem.txt")
//	defer file.Close()
//	checkErr(err)
//	file.WriteString(filename)
//}
func usage(){
	fmt.Printf(`
Usage:
  add:   添加信息
  list:  列出所有用户信息
  save:  保存用户信息到StudentSystem.txt
  exit:  退出系统
  help:  查看帮助
`)
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
			// add 1 Tom
			var id int
			var name string
			fmt.Sscan(line, &cmd, &id, &name)
			students[name] = Student{Id: id, Name: name}
		//case "del":
		case "list":
			fmt.Println("ID\tNAME")
			for _, v := range students {
				fmt.Printf("%v\t%v\n", v.Id, v.Name)
			}
		case "load":
			// read from file
			// json.UnMarshal -> students
			filename,err := ioutil.ReadFile("StudentSystem.txt")
			checkErr(err)
			json.Unmarshal(filename,&students)
			fmt.Println("ID\tName")
			//不是按照顺序的,是随机的
			for _,v := range students{
				fmt.Printf("%d\t%s\n", v.Id, v.Name)
			}
		case "save":
			// json.Marshal  students -> string
			// write to file
			bufFile, _ := json.Marshal(students)
			fmt.Printf("%v\n", string(bufFile))
			file, err := os.Create("StudentSystem.txt")
			defer file.Close()
			checkErr(err)
			file.WriteString(string(bufFile))
		case "help":
			// print help message
			usage()
		case "exit":
			// exit
			//0表示正常退出
			os.Exit(0)
		default:
			// do some stuff..
		}
	}
}
