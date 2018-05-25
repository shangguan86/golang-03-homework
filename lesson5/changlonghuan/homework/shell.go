package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	host, _ := os.Hostname()
	/*
	//想获取用户,发现string()之后,默认会增加换行
	checkUser := exec.Command("whoami")
	checkResult, err := checkUser.CombinedOutput()
	checkErr(err)
	user := string(checkResult)
	fmt.Printf(user)
	*/
	prompt := fmt.Sprintf("[longhuan@%s]$ ", host)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(prompt)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		args := strings.Fields(line)
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		checkErr(err)
		//fmt.Println()
		fmt.Printf(prompt)
	}
}
