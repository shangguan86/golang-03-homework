package main

import "fmt"

func main() {

	sss := "aaa中国我爱你"

	arr := []byte(sss)

	fmt.Println(arr[0])
	//fmt.Printf("arr type  is a %v\n", reflect.TypeOf(arr))
	arr[0] = 'x'
	fmt.Println(arr[0])
	fmt.Println(string(arr))

	n := 2
	if n > 2 {

		fmt.Println("big")

	} else if n < 2 {
		fmt.Println("small")

	} else {

		fmt.Println("other")
	}

}
