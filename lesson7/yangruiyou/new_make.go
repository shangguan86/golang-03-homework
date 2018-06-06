package main

import "fmt"

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func main() {
	//ok
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1

	fmt.Println(*y)

	//not ok,cannot make type Bar
	//z := make(Bar)
	//(*z).thingOne = "kkgo"
	//(*z).thingTwo = 2
	//fmt.Println(*z)

	//ok
	x:=make(Foo)
	x["x"]="goodbye"
	x["y"]="hello"

	x["x"]="goodbye1"
	x["y"]="hello1"

	fmt.Println(x)


}
