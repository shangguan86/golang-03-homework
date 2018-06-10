package main

import (
	"fmt"
)

func main() {

	//fmt.Println(utf8.RuneCountInString("美丽的中国china"))

	f := NewFile(2, "2.txt")
	fmt.Println(*f)
}

type File struct {
	fd   int
	name string
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}
