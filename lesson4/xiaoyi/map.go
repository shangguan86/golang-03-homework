package main

import "fmt"

func main() {
	// 类似于key/value格式,map[keytype]valuetype
	ages := make(map[string]int) // 声明方式1
	ages["a"] = 1
	ages["b"] = 2

	ages1 := map[string]int { // 声明方式2
		"a":1,
		"b":2,
	}
	ages1["c"] = 3
	delete(ages1,"b")

	fmt.Println(ages,ages1)
	//获取map元素
	fmt.Println(ages["a"],ages1["a"])
	//更新ages中a的值
	ages["a"] = ages["b"] + 2
	fmt.Println(ages,ages1)
	//判断map中是否存在某个key
	c,ok := ages["c"]
	if ok {
		fmt.Println(c)
	}else {
		fmt.Println("not found")
	}
	// or
	if c,ok := ages["c"];ok {
		fmt.Println(c)
	}
}
