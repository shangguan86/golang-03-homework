package main

import  (
    "fmt"
    "log"
    "os"
)

func main() {
    f, _ := os.Open("./")
    infos, err := f.Readdir(-2)
    if err != nil {
        log.Fatal(err) 
    }
    
    for _, info := range infos {
        flag := "d"
        if ! info.IsDir() {
            flag = "-" 
        }
        fmt.Printf("%v %v %v\n", flag, info.Size(), info.Name()) 
    }
}
