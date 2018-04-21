package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []int{1, 2, 3}
	fmt.Println(reflect.TypeOf(a))
}
