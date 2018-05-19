package main

import (
	"os/exec"
	"log"
	"fmt"
)

func main() {
	cmd := exec.Command("ls", "-lh")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
