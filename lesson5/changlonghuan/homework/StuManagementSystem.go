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

func usage() {
	fmt.Printf(`
Usage:
  - add:     添加用户信息
      add ID NAME
        ID:   用户ID
        NAME: 用户名称
  - list:    列出所有用户信息
  - del:     删除用户信息
      del NAME
        NAME: 用户名称
  - save:    保存用户信息
      save FILENAME
        FILENAME: 文件名称
  - load:    从文件读取信息
      load FILENAME    
        FILENAME: 文件名称
  - updata:  更新用户信息
      updata ID NAME
        ID:   用户ID
        NAME: 用户名称
  - exit|quit:  退出系统
  - help:    查看帮助（显示此帮助信息）
`)
}
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Add(id int, name string) {
	bufName := []rune(name)
	//判断名称是否输入
	if len(name) != 0 {
		//判断name是否为英文，不知道怎么判断中文，--！
		for _, v := range bufName {
			if v >= 65 && v <= 90 || v >= 97 && v <= 122 {
				continue
			} else {
				fmt.Printf("==> 请输入正确的用户名: %s\n\t用户名不能包含数字和特殊字符\n", name)
				return
			}
		}
		if _, ok := students[name]; ok {
			fmt.Println("用户已存在")
		} else {
			students[name] = Student{Id: id, Name: name}
			fmt.Printf("==> 用户添加成功...\n\tid: %d\n\tname: %s\n", id, name)
		}
	}else{
		fmt.Println("==> 请输入正确的ID和NAME")
	}
}

func List(student map[string]Student) {
	fmt.Println("==> 用户信息\nID\tNAME")
	for _, v := range student {
		fmt.Printf("%v\t%v\n", v.Id, v.Name)
	}
}
func Del(student map[string]Student, name string) {
	if len(name) == 0 {
		fmt.Println("==> 请输入要删除的用户")
	}else if _, ok := student[name]; ok {
		fmt.Printf("==> 用户删除成功...\n\tname: %s\n", name)
		delete(student, name)
		//可不打印，如果用户比较多
		List(student)
	} else {
		fmt.Printf("==> 用户不存在:%s\n",name)
	}
}
func Save(student map[string]Student, Stufile string) {
	if len(Stufile) == 0 {
		fmt.Println("==> 请输入需要保存的文件名称")
	} else {
		bufFile, err := os.Create(Stufile)
		defer bufFile.Close()
		checkErr(err)
		bufData, _ := json.Marshal(student)
		bufFile.WriteString(string(bufData))
		fmt.Println("==> 保存到文件: ", Stufile)
	}
}
func Load(student map[string]Student, Stufile string) {
	if len(Stufile) == 0 {
		fmt.Println("==> 请输入需要读取的文件名称")
	} else {
		bufFile, err := ioutil.ReadFile(Stufile)
		checkErr(err)
		json.Unmarshal(bufFile, &student)
		//可不打印，如果用户比较多
		List(student)
	}
}
func Update(student map[string]Student, id int, name string) {
	if len(name) == 0 {
		fmt.Println("==> 请输入正确的ID和NAME")
	} else if _, ok := student[name]; ok {
		fmt.Printf("==> 用户更新成功...\n\tid: %d\n\tname: %s\n", id, name)
		students[name] = Student{Id: id, Name: name}
		List(student)
	} else {
		fmt.Println("==> 用户不存在,添加用户成功")
		Add(id, name)
		//可不打印，如果用户比较多
		List(student)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()
		var StuFile string
		var Cmd string
		var id int
		var name string
		fmt.Sscan(line, &Cmd)
		switch Cmd {
		case "add":
			// add 1 Tom
			fmt.Sscan(line, &Cmd, &id, &name)
			Add(id, name)
		case "del":
			fmt.Sscan(line, &Cmd, &name)
			Del(students, name)
		case "list":
			List(students)
		case "save":
			fmt.Sscan(line, &Cmd, &StuFile)
			Save(students, StuFile)
		case "load":
			fmt.Sscan(line, &Cmd, &StuFile)
			Load(students, StuFile)
		case "update":
			fmt.Sscan(line, &Cmd, &id, &name)
			Update(students, id, name)
		case "help":
			usage()
			//测试多值
		case "exit", "quit":
			// exit
			//0表示正常退出
			os.Exit(0)
		default:
			// do some stuff..
			usage()
		}
	}
}