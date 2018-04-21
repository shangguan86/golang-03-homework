package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("input invaild args")
		return
	}

	inputSteps,err:=strconv.ParseInt(os.Args[1],10, 62)

	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(inputSteps)

	//if os.Args[1] == "1" {
	//	fmt.Println('1')
	//}
}

func step() {

}
