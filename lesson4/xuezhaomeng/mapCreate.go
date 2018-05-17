package main

import "fmt"

func main() {
	// 1.
	ages := make(map[string]int)
	ages["a"] = 1
	ages["b"] = 2

	for k, v := range ages {
		fmt.Println(k, v)
	}

	// 2.
	ages1 := map[string]int{"a": 1}
	ages1["b"] = 2

	for k, v := range ages1 {

		fmt.Println(k, v)
	}

	// 删除原属
	delete(ages1, "e")
fmt.Println(ages1)

	for k, v := range ages1 {

		fmt.Println(k, v)
	}

	// 获取值
	//1
	fmt.Println(ages["a"])
	//2
	c, ok := ages["c"]
	if ok {
		fmt.Println(c)
	} else {
		fmt.Println("not found")
	}

	//3
	if v, ok := ages["a"]; ok {
		fmt.Println(v)

	}

}
