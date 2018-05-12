package main

import (
    "crypto/md5"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    s := "abc"
    fmt.Printf("%v\n", []byte(s))

    filename := os.Args[1] 
    f, _ := ioutil.ReadFile(filename)
    md5sum := md5.Sum(f)
    fmt.Printf("md5: %x, len:%d\n", md5sum, len(md5sum))
}
