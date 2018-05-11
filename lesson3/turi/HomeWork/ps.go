package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func getname(proc string) string {
	filepath := "/proc/" + proc + "/cmdline"
	name, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return string(name)
}

func getpid(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	infos, err := f.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	for _, info := range infos {
		if info.IsDir() {
			s, err := strconv.Atoi(info.Name())
			if err != nil {
				continue
			}
			name := getname(info.Name())
			fmt.Printf("%v\t%v\n", s, name)
		}
	}
}

func main() {
	getpid("/proc")
}
