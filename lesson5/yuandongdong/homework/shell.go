package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

// TODO: homework
func main() {

	host, _ := os.Hostname()
	user, _ := user.Current()
	username := user.Username
	prompt := fmt.Sprintf("[%s:myshell@%s]$ ", username, host)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print(prompt)
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if line == "exit" {
			fmt.Println("exit my shell")
			os.Exit(1)
		}
		//todo : chdir()

		args := strings.Fields(line)
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		fmt.Print(prompt)
		if err != nil {
			fmt.Println(err)
			fmt.Print(prompt)
		}
	}

}
