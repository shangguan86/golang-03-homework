package menu

import (
	"bufio"
	"fmt"
	"os"

	stu "student_manager/student_manager"
)

func Menu(students *stu.Students, default_file_name string) {

	fmt.Println("===============学生管理系统==============")
	var err error
	var id int
	var stu_name, old_stu_name, new_stu_name, file_name string

	stu := stu.Student{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$>")
		scanner.Scan()
		line := scanner.Text()
		var cmd string
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "add":
			fmt.Sscan(line, &cmd, &id, &stu_name)
			stu.Id = id
			stu.Name = stu_name
			err = students.Add(stu)
			print_error(err)
		case "update":
			fmt.Sscan(line, &cmd, &old_stu_name, &new_stu_name)
			err = students.Update(old_stu_name, new_stu_name)
			print_error(err)
		case "del":
			fmt.Sscan(line, &cmd, &stu_name)
			err = students.Delete(stu_name)
			print_error(err)
		case "list":
			students.List()
		case "save":
			fmt.Sscan(line, &cmd, &file_name)
			err = students.Save(file_name, default_file_name)
			print_error(err)
		case "load":
			fmt.Sscan(line, &cmd, &file_name)
			err = students.Load(file_name, default_file_name)
			print_error(err)
		case "exit":
			os.Exit(0)
		default:
			students.Help()
		}
	}
}

func print_error(err error) {
	if err != nil {

		fmt.Println(err)
	}
}
