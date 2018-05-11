package main

import (
	"os"
	"log"
	"fmt"
	"strconv"
)

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	infos, err := f.Readdir(-2)
	if err != nil {
		log.Fatal(err)
	}

	for _, info := range infos {
		flag := "-"
		if info.IsDir() {
			flag = "d"
		}
		n, err := strconv.Atoi(info.Name())
		if err != nil {
			continue

		}

		if info.IsDir() {
			filename := path + "/" + info.Name() + "/cmdline"
			//作业:读取filename里面的命令,打印pid,cmd
			fmt.Println(filename)
			cmd := " "

		}

		fmt.Printf("%v %dB %s\n", flag, info.Size(), info.Name())

	}
	f.Close()

}
