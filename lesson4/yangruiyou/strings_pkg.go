package main

import (
	"fmt"
	"strings"
)

func main() {
	names := "没乐山|kkgo,kkgo,may here."
	for _, name := range strings.Split(names, ",") {
		fmt.Printf("%s \n ", name)
	}

	ManySpitString([]string{"kkgo*kkgo1*kkgo2|kkgo3|kkgo4\t\nkkgo999*kkgo6, kkgo3|kkgo4\tkkgo5*kkgo6",})

}

func ManySpitString(str []string) {
	for _, record := range str {
		fmt.Println(strings.FieldsFunc(record, func(char rune) bool {
			switch char {
			case '\t', '*', '|', '\n', ' ', ',':
				return true

			}
			return false
		}))
	}
}
