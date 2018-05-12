package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	// linux: /proc, 其他系统自己创建个模拟目录
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Args is Not")
		os.Exit(1)
	}
	path := os.Args[1]
	f, _ := os.Open(path)

	infos, err := f.Readdir(-2)
	if err != nil {
		log.Fatal(err)
	}

	for _, info := range infos {
		// 判断是否是pid
		_, err := strconv.Atoi(info.Name())
		if err != nil {
			continue
		}

		// 判断是否是pid目录
		if info.IsDir() {
			// 拼接类似/proc/1877/cmdline的路径
			filename := path + "/" + info.Name() + "/cmdline"
			fp, err := os.Open(filename)
			if err != nil {
				log.Fatal(err)
			}

			r := bufio.NewReader(fp)

			cmd := ""
			for {
				line, _, err := r.ReadLine()
				if err == io.EOF {
					break
				}
				cmd += string(line)
			}
			// 作业: 读取filename里的命令, 打印 pid, cmd
			//
			//fmt.Println(filename)
			if cmd != "" {
				fmt.Printf("%v\t%v\n", info.Name(), cmd)
			}
		}
	}
}
