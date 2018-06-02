package main

import "fmt"

type gopher struct {
	name string
	age int
}

func main() {
	//s := "hello大家好"
	//r := []rune(s)
	//
	//fmt.Println(len(s))
	//
	//arr := []byte(s)
	//arr[0] = 'w'
	//r[0] = 'y'
	//fmt.Printf("%v\n",s)
	//fmt.Printf("%v\n",string(arr))
	//fmt.Printf("%v\n",string(r))

	//fmt.Println()
	//fmt.Printf("%v\n",reflect.TypeOf(r))
	//for i:=0;i<len(r);i++{
	//	fmt.Println(string(r[i]))
	//}

	//var name  = gopher{"wangfei0",0}
	//fmt.Printf("%v\n",name)
	//name.name ="wangfei1"
	//name.age = 1
	//fmt.Printf("%v\n",name)
	//name.name,name.age= "wangfei2",2
	//fmt.Printf("%v\n",name)
	//
	//s1 := make([]int,0)
	//s1 = append(s1,1,2,3)
	//fmt.Printf("%v\n",s1)
	//fmt.Printf("%v\n",reflect.TypeOf(s1))


	//var arr1 [5]int{}
	//arr1 := [5]int{}
	//arr1 := [...]int{1,2,3,4,5,6}
	//fmt.Printf("%v\n",arr1)

	call()
}

func call() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer func() { fmt.Printf("1\n") }()
	defer func() { fmt.Printf("2\n") }()
	defer func() { fmt.Printf("3\n") }()
	panic("bone")
}
