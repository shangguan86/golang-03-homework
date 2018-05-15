package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 添加
func add(id int, name string) {
	if len(name) < 1 {
		fmt.Println("必须输入学生名字")
		return
	}
	if id == 0 {
		fmt.Println("必须输入学生id, 不能为0")
		return
	}

	if _, ok := students[name]; ok {
		fmt.Printf("%s 学生信息已存在\n", name)
		return
	}

	students[name] = Student{
		Id: id, Name: name,
	}
	fmt.Println("添加成功")
}

// 列表
func list() {
	fmt.Printf("Id\tName\n")
	for _, v := range students {
		fmt.Printf("%d\t%s\n", v.Id, v.Name)
	}
}

// 更新
func update(old_stu_name, new_stu_name string) {
	info, ok := students[old_stu_name]
	if !ok {
		fmt.Printf("%s 学生信息不存在\n", old_stu_name)
		return
	}
	if len(new_stu_name) < 1 {
		fmt.Println("必须输入新名字")
		return
	}
	info.Name = new_stu_name
	delete(students, old_stu_name)
	students[new_stu_name] = info
	fmt.Printf("更新成功:%v\n", info)

}

// 删除
func del(stu_name string) {
	if len(stu_name) < 1 {
		fmt.Println("必须输入学生名字")
		return
	}
	delete(students, stu_name)
	fmt.Printf("删除:%s 信息成功\n", stu_name)

}

// 保存
func save(file_path string) {
	// 获取文件路径
	fpath := get_file_path(file_path)

	// 解析成json格式
	buf, err := json.MarshalIndent(students, "", "\t")
	if err != nil {
		fmt.Println("解析成json格式:", err)
		return
	}

	// 覆盖写文件,ioutil自动关闭文件句柄
	err = ioutil.WriteFile(fpath, buf, 0600)

	if err != nil {
		fmt.Println("写文件失败:", err)
		return
	}
}

// 加载
func load(file_path string) {
	// 获取文件路径
	fpath := get_file_path(file_path)

	// 读取文件
	bytes, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Println("读取文件错误:", err.Error())
		return
	}

	// 反解析json 到 struct
	if err = json.Unmarshal(bytes, &students); err != nil {
		fmt.Println("反解析json文件失败", err.Error())
		return
	}
	fmt.Println("加载学生信息成功")

}

// 获取文件路径
func get_file_path(file_path string) string {

	// 获取文件路径,如果返回异常则加载默认的文件
	_, err := os.Stat(file_path)
	var fpath string
	if err != nil {
		fpath, err = filepath.Abs(filepath.Dir("."))
		fpath = filepath.Join(fpath, "stu.json")
		fmt.Printf("%v,文件路径不存在使用默认路径:%s\n", err, fpath)
	}
	return fpath
}
