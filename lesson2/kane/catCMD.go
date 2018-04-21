package main

import (
	"io/ioutil"
	"fmt"
	"flag"
)
var filePath string

func init() {
	flag.StringVar( &filePath,"f", "/", "Please input a file path")
}
func main()  {
	flag.Parse()
	printFile(filePath)

}

func printFile(name string)  {

	buf,err := ioutil.ReadFile(name)

	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(buf)


}
