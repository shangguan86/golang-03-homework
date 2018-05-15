package main

import (
	"fmt"
	"strings"
)

// 方法1
func count_word_method1(words []string) {

	count_words := make(map[string]int)

	for _, value := range words {
		count_words[value]++
	}
	fmt.Printf("方法1 单词总数:%v\n", count_words)

}

// 方法2
func count_word_method2(words []string) {

	count_words := make(map[string]int)

	for _, key := range words {
		// 判断map中是否有这个key,有则加1,否则初始化此key的值为1
		if _, ok := count_words[key]; ok {
			count_words[key] = 1
		} else {
			count_words[key] = count_words[key] + 1
		}

	}
	fmt.Printf("方法2 单词总数:%v\n", count_words)

}

func main() {
	s := "one two three one"

	words := strings.Fields(s)

	count_word_method1(words)
	count_word_method2(words)
}
