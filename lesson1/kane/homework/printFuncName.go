package main

import (
	"runtime"
	"fmt"
	"strings"
)

func main() {

	pc, _, _, ok := runtime.Caller(0)

	if ok {
		funcForPc := runtime.FuncForPC(pc)
		funcName := strings.Split(funcForPc.Name(),".")
		fmt.Println(funcName[1])
	}

}

func Test()  {
	pc, _, _, ok := runtime.Caller(0)

	if ok {
		funcForPc := runtime.FuncForPC(pc)
		funcName := strings.Split(funcForPc.Name(),".")
		fmt.Println(funcName[1])
	}
}