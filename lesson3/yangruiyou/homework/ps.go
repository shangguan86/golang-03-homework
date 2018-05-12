package main

import (
	"os"
	"log"
	"fmt"
	"strconv"
	"io/ioutil"
)

func main() {
	path := os.Args[1]

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	//f.Close()

	infos, err := f.Readdir(-2)
	if err != nil {
		log.Fatal(err)
	}

	for _, info := range infos {
		flag := "-"
		if info.IsDir() {
			flag = "d"
		}
		_, err := strconv.Atoi(info.Name())
		if err != nil {
			continue

		}

		if info.IsDir() {
			filename := path + "/" + info.Name() + "/cmdline"
			//作业:读取filename里面的命令,打印pid,cmd

			f, _ := ioutil.ReadFile(filename)
			cmd := string(f)
			fmt.Printf("%v %dB %s\n", flag, info.Size(), cmd)

		}

	}
	f.Close()

}
