package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	S1 := "hello" + "word"
	fmt.Println(S1)
	//打印下标位0 的字符，显示的为 ASCII 码
	fmt.Println(S1[0])
	//打印下标位0 的字符，转化为对应的字符
	fmt.Println(string(S1[0]))
	//打印 S1字符的长度
	fmt.Println(len(S1))
	//获取字符串的切片， 切片左闭右开
	fmt.Println(S1[0:3])
	fmt.Println(S1[0:len(S1)])
	//for i := 0; i < len(S1); i++ {
	//	fmt.Println(S1[i], string(S1[i]))

	//}

	//将字符串 string转为为 int 类型
	i, _ := strconv.Atoi("-42")
	//将int 转为为 string 类型
	fmt.Printf("i type  is a %v\n", reflect.TypeOf(i))

	//打印变量的类型
	s := strconv.Itoa(-42)
	fmt.Printf("s type  is a %v\n", reflect.TypeOf(s))

	f, _ := strconv.ParseFloat("3.14", 32)
	fmt.Printf("f(32) is a %v\n", f)
	fmt.Printf("f(32) type  is a %v\n", reflect.TypeOf(f))
	f6, _ := strconv.ParseFloat("3.14", 64)
	fmt.Printf("f6(64)  is a %v\n", f6)
	fmt.Printf("f6(64) type  is a %v\n", reflect.TypeOf(f6))

	//中文转换
	sss := "aaa中国我爱你"
	//sss[0] = 'x', 报错，字符串为常量
	fmt.Println(sss)

	arr := []byte(sss)
	for k, v := range arr {

		fmt.Println(k, v)
	}

	//fmt.Printf("arr type  is a %v\n", reflect.TypeOf(arr))
	arr[0] = 'x'

	fmt.Println(string(arr))
}
