package main

import (
	"fmt"
	"os"
	"bufio"
	"time"
)

func main() {
	filename := "./log"
	s := time.Now()
	ReadString(filename)
	e1 := time.Now()
	fmt.Printf("ReadString:%v\n", e1.Sub(s))
	ReadLine(filename)
	e2 := time.Now()
	fmt.Printf("readLine:%v\n", e2.Sub(e1))


	fmt.Println("___________________")
	fmt.Println(20%8)
}

func ReadString(filename string) {
	f, _ := os.Open(filename)
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		_, err := r.ReadString('\n')
		if err != nil {
			break
		}
	}
}

func ReadLine(filename string) {
	f, _ := os.Open(filename)
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		_, err := readLine(r)
		if err != nil {
			break
		}
	}
}

func readLine(r *bufio.Reader) (string, error) {
	line, isprefix, err := r.ReadLine()
	for isprefix && err == nil {
		var bs []byte
		bs, isprefix, err = r.ReadLine()
		line = append(line, bs...)
	}
	return string(line), err

}
