package main 

// TODO: need import some packages

func printFile(name string) {
    buf, err := ioutil.ReadFile(name)
    if err != nil {
        fmt.Println(err) 
        return
    }
    fmt.Println(buf)
}


// TODO: flag.Parse()
// 读取命令行参数， 接受一个 -f 的参数， 代表文件名， 调用printFile()函数，打印文件内容

func main() {

}
