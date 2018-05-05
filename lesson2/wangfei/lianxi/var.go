package main

import "fmt"

func main(){
	//var i int = 10
	//var s string = "hello world1"
	//var f float64 = 1.1
	//fmt.Println(i,s,f)
	//
	//var i2 = 20
	//var s2 = "hello world2"
	//var f2 = 2.1
	//fmt.Println(i2,s2,f2)

	var (
		v1 int
		v2 string
		v3 float64
		v4 bool
	)

	v1 = 30
	v2 = "hello world3"
	v3 = 3.1
	v4 = true

	//fmt.Printf("%v,%v,%v\n",v1,v2,v3,v4)
	fmt.Printf("%d,%s,%.1f,%v\n",v1,v2,v3,v4)

	//s1 := []int{2,3,1,4,5,3}
	//fmt.Println(s1[1])

}
