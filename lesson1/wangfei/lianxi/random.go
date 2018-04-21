package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generate_random(num int)  {

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i:=1;i<=num;i++{
		rand_num := rand.Intn(100)
		fmt.Printf("随机数是:%d\n",rand_num)
	}
}

func get_random_name()  {
	lst_name := [...]string{"a","b","c"}
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand_num := rand.Intn(len(lst_name))
	fmt.Println(lst_name[rand_num])
}

func main() {
	generate_random(1)
	get_random_name()
}
