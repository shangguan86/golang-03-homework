package main

import "fmt"

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (iota * 10)
	MB
	GB
	TB
	PB
)

func main() {
	fmt.Println(KB, MB, GB, TB, PB)
}
