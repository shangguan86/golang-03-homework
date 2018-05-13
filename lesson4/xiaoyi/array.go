package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/*
	   认识数组
	*/
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	// 迭代数组值方式1
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// 迭代数组值方式2
	for i := range a {
		fmt.Printf("%d %d\n", i, a[i])
	}
	// 迭代数组值方式3
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	/*
	   初始化数组
	*/
	//创建数组方式1
	var r [3]int = [3]int{1, 2, 3}
	// int根据操作系统的位数决定其大小,因此int64/8*3=24;
	fmt.Printf("r:%v,len:%d,size:%v\n", r, len(r), unsafe.Sizeof(r)) // [1,2,3],3,24
	var b [3]int = [3]int{1, 2}
	//空缺位使用默认值0填充;
	fmt.Printf("b:%v,len:%d,size:%v\n", b, len(b), unsafe.Sizeof(b)) // [1,2,0],3,24

	//创建数组方式2
	c := [...]int{1, 2, 3}
	c[2] = 0                                                         //更改c数组index为2的值为0,也就是将3改为0;
	fmt.Printf("c:%v,len:%d,size:%v\n", c, len(c), unsafe.Sizeof(c)) // [1,2,0],3,24

	d := [...]int{0: 1, 2: 3}
	fmt.Printf("d:%v,len:%d,size:%v\n", d, len(d), unsafe.Sizeof(d)) // [1,0,3],3,24

	/*
	   数组变量赋值和比较
	*/
	a1 := [3]int{1, 2, 3}
	fmt.Println(a1)

	var a2 [3]int
	a2 = a1
	fmt.Printf("a1等于a2:%v\n", a1 == a2)
	//内存连续，8byte
	for i := 0; i < 3; i++ {
		fmt.Printf("a1[%d]:%d,&a1[%d]:%p,&a1:%p\n", i, a1[i], i, &a1[i], &a1)
	}
	for i := 0; i < 3; i++ {
		fmt.Printf("a2[%d]:%d,&a2[%d]:%p,&a2:%p\n", i, a2[i], i, &a2[i], &a2)
	}
}
