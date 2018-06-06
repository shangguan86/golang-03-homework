package main

import (
	"encoding/json"
	"log"
	"os"
)

type Files struct {
	Filename  string
	FileSize  int64
	Directory bool
}

type Sfiles []Files

var files Sfiles

func (p *Sfiles) Append(name string, size int64, dir bool) {

	f := Files{Filename: name, FileSize: size, Directory: dir}
	*p = append(*p, f)

}

func (p *Sfiles) Getinfo() (b []byte, err error) {

	b, err = json.Marshal(*p)
	return

}

func FileGet() (b []byte, err error) {
	return files.Getinfo()

}

func CheckError(err error) {

	if err != nil {
		log.Fatal(err)
	}

}

func ReadDir(dir *string) {

	files = make([]Files, 0)

	f, err := os.Open(*dir)
	CheckError(err)
	defer f.Close()

	infos, _ := f.Readdir(-1)
	for _, info := range infos {

		tag := false
		if info.IsDir() {
			tag = true
		}

		files.Append(info.Name(), info.Size(), tag)

	}

}
