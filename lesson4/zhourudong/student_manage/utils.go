package main

import (
	"bufio"
	"fmt"
	"os"
)

// 菜单
func menu() {
	fmt.Println("============ 学生管理系统 ================")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$>")
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
		case "update":
			var old_stu_name, new_stu_name string
			fmt.Sscan(line, &cmd, &old_stu_name, &new_stu_name)
			update(old_stu_name, new_stu_name)
		case "del":
			var stu_name string
			fmt.Sscan(line, &cmd, &stu_name)
			del(stu_name)
		case "list":
			list()
		case "save":
			var file_path string
			fmt.Sscan(line, &cmd, &file_path)
			save(file_path)
		case "load":
			var file_path string
			fmt.Sscan(line, &cmd, &file_path)
			load(file_path)

		case "exit":
			os.Exit(0)

		default:
			fmt.Println(help_text)
		}
	}
}

var help_text = `使用示例:

添加学生:
	$>add 1 张三

更新学生:
	$>update 张三 李四

列出学生信息:
	$> list 

删除学生:
	$>del 张三

保存学生信息到文件中:
	$>save [file_abs_path]
	or
	$>save 

加载:
	$>load [file_abs_path]
	or 
	$> load

退出:
	$>exit
`
