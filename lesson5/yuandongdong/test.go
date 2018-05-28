package main

import (
	"fmt"
	"os"
)

func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("[user@%s]$ ", host)
	fmt.Println(prompt, host)

}
