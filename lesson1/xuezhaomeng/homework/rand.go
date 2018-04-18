package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Intn(100))
}
