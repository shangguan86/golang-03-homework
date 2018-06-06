package main

import (
	"bufio"
	"os"
	"fmt"
)

func main(){


	inputReader:=bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")
	input,err:=inputReader.ReadString('\n')
	if err!=nil{
		fmt.Printf("an error occurred:%s\n",err)
		os.Exit(1)
	}else{
		name:=input[:len(input)-1]
		fmt.Printf("hello,%s,what can I do for you?\n",name)
	}
}