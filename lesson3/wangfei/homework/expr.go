package main

import (
	"os"
	"strconv"
	"fmt"
	"log"
)

func main(){
s1 := os.Args[1]
op := os.Args[2]
s2 := os.Args[3]

//string转换成Int
n1,err := strconv.Atoi(s1)
//fmt.Println(reflect.TypeOf(n1))
if err !=nil{
log.Fatal(err)
}

n2,err := strconv.Atoi(s2)
if err !=nil{
log.Fatal(err)
}

switch op {
case "+":
fmt.Printf("%d\n",n1 + n2)
case "-":
	fmt.Printf("%d\n",n1 - n2)
case "*":
	fmt.Printf("%d\n",n1 * n2)
case "/":
	//go run expr.go 2 \* 2,需要转义
	fmt.Printf("%d\n",n1 / n2)
default:
fmt.Println("error")
}
}

