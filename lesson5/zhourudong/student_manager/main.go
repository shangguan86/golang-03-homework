package main

import (
	"flag"

	m "student_manager/menu"
	. "student_manager/student_manager"
)

var default_file_name string

func init() {
	default_file_name = "stu.json"
	flag.StringVar(&default_file_name, "default_file_name", default_file_name, "默认保存文件名")
	flag.Parse()
}

func main() {
	students := &Students{}
	m.Menu(students, default_file_name)
}
