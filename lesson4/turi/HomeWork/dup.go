package main

import "fmt"

func remove(s []string, i int) []string {
	copy(s[i:], s[i+1:])
	//s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func dup(s []string) []string {
	n := len(s)
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			n -= 1
		}
		if n+i > len(s) {
			break
		}
	}
	return s[:n]

}

func main() {
	s := []string{"Tom", "Tom", "Make", "Make", "Tom", "JOM", "Make", "Tom", "Tom", "Tom", "Tom", "Tom", "Jack", "test"}
	// s := []string{"tom", "tom", "make", "make", "mom", "tom", "tom"}
	fmt.Println(dup(s))

}
