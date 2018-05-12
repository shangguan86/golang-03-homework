package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

//按行读取
func printFile(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
		return
	}
//记得关闭文件句柄
defer f.Close()
	r := bufio.NewReader(f)
	for {
		buf, _, err := r.ReadLine()
	//	if err == io.EOF {
	//		return
	//	} else if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
      if err != nil {
            break
        }
		fmt.Println(string(buf))

	}
}

//解析-f参数
func Parse() string {
	fname := flag.String("f", "", "file name")
	flag.Parse()
	if fname == nil {
		fmt.Println("no argument")
	}
	return *fname
}

func main() {
	f := Parse()
	printFile(f)

}
