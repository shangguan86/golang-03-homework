package main

import (
	"os"
	"log"
	"fmt"
	"strconv"
	"bufio"
)

func read_file(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	line, err := r.ReadString('\n')
	f.Close()
	return string(line)
}

func read_dir() {
	path := "/proc"
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	infos, _ := f.Readdir(-1)
	for _, v := range infos {
		if v.IsDir() {
			_, err := strconv.Atoi(v.Name())
			if err == nil {
				pid := v.Name()
				filename := path + "/" + v.Name() + "/cmdline"
				cmd_line := read_file(filename)
				fmt.Println(pid, cmd_line)
			}
		}
	}
}

func main() {
	read_dir()
}
