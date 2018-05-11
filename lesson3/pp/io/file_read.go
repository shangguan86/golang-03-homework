package main


import (
    "fmt"
    "os"
    "io"
    "log"
    "bufio"
    "reflect"
)

func main() {
    filename := os.Args[1]
    f, err := os.Open(filename)
    fmt.Println("type of f: ", reflect.TypeOf(f))
    if err != nil {
        log.Fatal(err) 
    }
    r := bufio.NewReader(f)

    for {
        // \n为读取的分隔符
        line, err := r.ReadString(' ') 
        if err == io.EOF {
            break 
        }
        fmt.Print(line)
    }

    f.Close()
}
