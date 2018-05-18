package main

import (
	"fmt"
	"strings"
)

//## 作业1:
//统计单词出现的个数(使用Map), 提示代码：

//```go
//s := "one two three one"
//var words []string
//words = strings.Fields(s)
//```

func compute_words_count(str string) {
	var words []string
	words = strings.Fields(str)
	count := make(map[string]int)
	for _,value := range words{
			count[value] += 1
	}
	fmt.Println(count)
}

func main() {
	s := "one two three one"
	fmt.Printf("string is :%s\n", s)
	compute_words_count(s)
}
