package main

import (
	"strings"
	"fmt"
)

// 统计单词出现的个数(使用Map), 提示代码：

func CountWord(s string) map[string]int {
	var strs map[string]int = make(map[string]int)
	fmt.Println(strs)
	data := strings.Fields(s)
	for _, str := range data {
		index, ok := strs[str]
		if ok {
			index ++
			strs[str] = index
		} else {
			strs[str] = 1
		}
	}
	return strs
}

func main() {

	fmt.Println(CountWord("one two three one"))

}
