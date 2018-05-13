// 对slice的操作，包括创建/添加/删除等;
package main

import "fmt"

func PrintSlice(e string,f []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", e, len(f), cap(f), f)
}

func RemoveSlice(s []int,i int) []int {
	copy(s[i:],s[i+1:]) // 保持slice顺序不变
	//s[i] = s[len(s)-1] // slice 顺序将发生改变
	return s[:len(s)-1]
}

func main() {
	//make 创建slice
	A := make([]int,5) // 创建slice A并且len为5
	PrintSlice("A",A)

	B := make([]int,0,5) //创建slice B并且指定len为0,容量cap为5;
	PrintSlice("B",B)

	C := B[:2]
	PrintSlice("C",C)

	D := C[2:5]
	PrintSlice("D",D)

	//append slice,appen slice时,如果超出当前slice容量就会创建一个容量为 2*cap(slice)新的slice;
	var S []int
	fmt.Printf("len=%d cap=%d %v\n", len(S), cap(S), S)

	S = append(S,0)
	fmt.Printf("len=%d cap=%d %v\n", len(S), cap(S), S)

	S = append(S,1)
	fmt.Printf("len=%d cap=%d %v\n", len(S), cap(S), S)

	S = append(S,2,3,4)
	fmt.Printf("len=%d cap=%d %v\n", len(S), cap(S), S)

	s := []int{1,2,3,4}
	fmt.Println(RemoveSlice(s,1))
}
