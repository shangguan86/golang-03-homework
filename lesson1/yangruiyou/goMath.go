package main

import (
	"fmt"
	"math"
)

func main(){
	a:=-2.0

	fmt.Println(math.Abs(a))//2

	b:=math.Pow(a,2)
	fmt.Println(b)//4

	fmt.Println(math.Sqrt(b))//2
}
