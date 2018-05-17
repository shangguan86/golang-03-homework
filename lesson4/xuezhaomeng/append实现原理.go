package main

import "fmt"

func main() {
	s := make([]int, 0, 3)
	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("cap %d , len %d , s:%p , value %v\n", cap(s), len(s), s, s)
		//  %p 查看地址

		//  当超过切片的容量时，会进行扩容，申请一个新的地址 ，为原容量的2倍
	}

}
