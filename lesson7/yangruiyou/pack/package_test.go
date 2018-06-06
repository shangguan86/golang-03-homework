package main

import (
	"fmt"
	"github.com/51reboot/golang-03-homework/lesson7/yangruiyou/pack"
)

func main() {
	var test1 string
	test1 = pack.ReturnStr()

	fmt.Printf("ReturnStr from package1:%s\n", test1)

}
