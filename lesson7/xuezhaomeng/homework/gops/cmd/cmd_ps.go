package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"time"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}

}

func WriteFile(b []byte) {
	err := ioutil.WriteFile("ps.txt", b, 0644)
	CheckError(err)
}

func Gops() {
	for {
		cmd := exec.Command("ps", "-ef")
		stdout, err := cmd.StdoutPipe()
		CheckError(err)
		defer stdout.Close()

		err = cmd.Start()
		CheckError(err)

		bs, err1 := ioutil.ReadAll(stdout)
		CheckError(err1)

		fmt.Println(string(bs))
		WriteFile(bs)

		time.Sleep(time.Second * 10)
	}

}
