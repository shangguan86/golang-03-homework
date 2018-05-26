package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"os/exec"
)

func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("[test_shell@%s]$ ", host)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(prompt)

	for scanner.Scan() {
		/*
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, os.Interrupt)
		*/
		fmt.Print(prompt)
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if line == "exit" {
			fmt.Println("退出shell")
			os.Exit(0)
		}
		args := strings.Fields(line)

		// cmd := exec.Command("/bin/bash")
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin

		// 处理 键盘各种中断信号 Ctrl + C , Ctrl + D ,
		/*		for {
					sig := <-sigs
					switch sig {

					case syscall.SIGTERM:
						os.Exit(-1)
					case syscall.SIGQUIT:
						os.Exit(-1)
					default:
						fmt.Printf("111 %#v %T", sigs, sigs)
					}
				}
		*/

		cmd.Stdout = os.Stdout
		fmt.Println()
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(prompt)
	}
}
