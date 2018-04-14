package main

import (
	"math/rand"
	"fmt"
	"strconv"
)

func main() {

	for j := 1; j <= 5; j++ {
		randNumber := rand.Intn(100)
		fmt.Println("The number is: " + strconv.Itoa(randNumber))
	}

}
