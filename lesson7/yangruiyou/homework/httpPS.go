package main

import (
	"os/exec"
	"bufio"
	"fmt"
)

func main() {
	execCommand()
	//http.HandleFunc("/ps", execPS)
	//http.ListenAndServe(":9000", nil)

}

//func execPS(w *http.ResponseWriter, req *http.Request) {
//	//
//
//}

func execCommand()string{
	cmd := exec.Command("/bin/bash", "-c", `ps`)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Printf("%s\n", err)
		// return
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("%s\n", err)
		// return
	}

	outPutBuf := bufio.NewReader(stdout)
	for {
		output, _, err := outPutBuf.ReadLine()
		if err != nil {
			fmt.Printf("%s\n", err)
			// return
		}

		//fmt.Printf("%s\n", string(output))
		return string(output)
	}
	if err := cmd.Wait(); err != nil {
		fmt.Printf("%s\n", err)
		// return
	}
}
