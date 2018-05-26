package student_manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type StudenterInterface interface {
	Add(student Student) error
	Update(old_stu_name, new_stu_name string) error
	Delete(student_name string) error
	List()
	Save(file_name, default_file_name string) error
	Load(file_name, default_file_name string) error
	Help()
}

type Student struct {
	Id   int    `json:"student_id"`
	Name string `json:"student_name" `
}
type Students map[string]Student

func (stus Students) Add(stu Student) error {
	if stu.Id == 0 || stu.Id < 0 {
		return errors.New("请输入学生id,且不能小于0")
	}
	if len(stu.Name) < 1 {
		return errors.New("请输入学生名字")
	}
	stus[stu.Name] = stu

	return nil
}

func (stus Students) Update(old_stu_name, new_stu_name string) error {
	info, ok := stus[old_stu_name]
	if !ok {
		return fmt.Errorf("%s 学生信息不存在", old_stu_name)
	}
	if len(new_stu_name) < 1 {

		return errors.New("必须输入新名字")
	}
	info.Name = new_stu_name
	// todo
	//delete(stus, old_stu_name)
	stus[new_stu_name] = info
	return nil
}

func (stus Students) Delete(student_name string) error {
	if len(student_name) < 1 {
		return errors.New("学生信息不存在\n")
	}
	delete(stus, student_name)
	return nil
}

func (stus Students) List() {
	fmt.Printf("Id\tName\n")
	for _, v := range stus {
		fmt.Printf("%d\t%s\n", v.Id, v.Name)
	}
	fmt.Print("$>")
}

func (stu Students) Save(file_name, default_file_name string) error {
	// 获取文件路径
	fpath := get_file_path(file_name, default_file_name)
	// 解析成json格式
	buf, err := json.MarshalIndent(stu, "", "\t")
	if err != nil {
		return fmt.Errorf("解析成json格式失败:%s\n", err)
	}

	// 覆盖写文件,ioutil自动关闭文件句柄
	err = ioutil.WriteFile(fpath, buf, 0600)

	if err != nil {
		return fmt.Errorf("写文件失败:%s\n", err)
	}
	return nil
}

func (stu Students) Load(file_name, default_file_name string) error {
	// 获取文件路径
	fpath := get_file_path(file_name, default_file_name)

	// 读取文件
	bytes, err := ioutil.ReadFile(fpath)
	if err != nil {
		return fmt.Errorf("读取文件失败:%s\n", err)
	}

	// 反解析json 到 struct
	if err = json.Unmarshal(bytes, &stu); err != nil {
		return fmt.Errorf("反解析json文件失败:%s\n", err)
	}
	fmt.Println("加载学生信息成功")
	return nil
}

func (stu Students) Help() {
	help_text := `
使用示例:

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
	fmt.Println(help_text)
}

// 获取文件路径
func get_file_path(file_path, default_file_name string) string {

	// 获取文件路径,如果返回异常则加载默认的文件
	_, err := os.Stat(file_path)
	var fpath string
	if err != nil {
		fpath, err = filepath.Abs(filepath.Dir("."))
		fpath = filepath.Join(fpath, default_file_name)
		fmt.Printf("文件路径不存在使用默认路径:%s %s\n", err, fpath)
		return fpath
	}
	return file_path
}
