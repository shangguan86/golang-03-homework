package main

import (
	"os"
	"bufio"
	"strings"
	"os/exec"
	"fmt"
	"log"
)

// TODO: homework

func main() {
	connect()

}

func connect() {
	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
		return
	}
	prompt := fmt.Sprintf("[pp@%s]$ ", host)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print(prompt)
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		args := strings.Fields(line)
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
