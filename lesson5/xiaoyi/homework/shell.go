package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"os/exec"
)

func main() {
	host,_ := os.Hostname() //获取主机名
	prompt := fmt.Sprintf("[xiaoyi@%s]",host) //将参数列表"host"填写到格式化字符串"[xiaoyi@%s]"中的占位符;
	scanne := bufio.NewScanner(os.Stdin)

	for  {
		fmt.Print(prompt)
		scanne.Scan()
		line := scanne.Text()

		if len(line) == 0 {
			continue
		}

		args := strings.Fields(line)
		cmd := exec.Command(args[0],args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}

		if line == "exit" {
			break
		}
	}
}