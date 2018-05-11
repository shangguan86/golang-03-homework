package main

import (
	"fmt"
	"reflect"
)

func main() {
	s1 := "hello中国"

	arr := []byte(s1)
	fmt.Println(arr)

	arr[0] = 'x'
	fmt.Println(s1, string(arr))

	for k, v := range arr {
		fmt.Println(k, v)
	}

	a := []byte(s1)
	fmt.Println(reflect.TypeOf(a))

	for k, v := range a {
		fmt.Println(k, string(v))
	}

	//a[5] = '美'
	//s2 := string(a)
	//fmt.Println(s2)
}
