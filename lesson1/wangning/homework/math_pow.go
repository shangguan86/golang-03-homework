package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	args := os.Args
	if args == nil || len(args) < 2 {
		log.Fatal("Please Input a Args")
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	num_pow := math.Pow(float64(num), 2)
	fmt.Printf("Number:%d math Pow:%.1f\n", num, num_pow)

}
