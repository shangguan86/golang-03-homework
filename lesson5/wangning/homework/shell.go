package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// TODO: homework

func main() {

	host, _ := os.Hostname()
	prompt := fmt.Sprintf("[shen@%s]$ ", host)
	scanner := bufio.NewScanner(os.Stdin)
	for  {
		fmt.Printf("%s", prompt)
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		args :=strings.Fields(line)
		if args[0] == "exit" {
			break
		}
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}

}
