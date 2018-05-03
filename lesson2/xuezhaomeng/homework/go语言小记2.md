### <center> go语言小记2 </center>



##### go通过package组织
- package关键字放在程序的第一行
- 两种package，一种是库package，一种是二进制package
- 二进制package使用main来表示，库package的名字跟go文件所在的目录名一样


##### 格式化代码&自动import ,-w保存
```
// 格式化代码
gofmt -w  hello.go

//自动import代码
goimport -w hello.go 
```


##### go for的表现形式
> Go语言中的循环语句只支持for关键字，而不支持while和do-while结构。

```
//循环有限的次数
sum := 0 
for i := 0; i < 10; i++ { 
    sum += i 
}  
```

```
//无限循环的写法：

sum := 0 
for { 
    sum++  
    if sum > 100 { 
        break 
    } 
} 
```

##### go 声明变量
* Go中变量的声明，使用var关键字
* 对声明却未使用的变量报错

```
//声明并初始化一个变量
var m int = 10

//声明初始化多个变量
var i, j, k = 1, 2, 3

//多个变量的声明(注意小括号的使用)
var(
   num  int
   name string
)

//快速声明，短变量
str := "Hello"

//丢弃变量 ,_表示
_: =  3


//变量在定义时没有明确的初始化时会赋值为 零值 。
//数值类型为 0 ，
//布尔类型为 false ，
//字符串为 "" （空字符串）。

```


##### 字符占位
> 由golang 的fmt 包实现

```
//普通占位符
%v      相应值的默认格式。

//布尔占位符
%t          true 或 false。

//整数占位符
%d      十进制表示   
%b      二进制表示   
%o      八进制表示   
%x      十六进制表示，字母形式为小写 a-f
%X      十六进制表示，字母形式为大写 A-F

//浮点数
.3%f      fmt.Printf("%.3f\n", y)保留小数点后三位
%f        fmt.Printf("%f", 10.2111) 默认小数点后六位

//字符串与字节切片
%s      输出字符串（string类型或[]byte) 
        fmt.Printf("%s\n", []byte("Go语言"))
        fmt.Printf("%s\n", "1111")


```

##### 指针
> 一个指针变量指向了一个值的内存地址。

###### 定义指针变量
```
var var_name *var-type
//var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针。
```

###### 为指针变量赋值
```
var a int= 20   /* 声明实际变量 */
var ip *int        /* 声明int类型指针变量 */
ip = &a  /* 将a变量的地址 赋值给ip指针变量 */
```

###### 访问指针变量中指向地址的值
> 在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。
```
fmt.Printf("a 变量的地址是: %x\n", &a  )

/* 指针变量的存储地址 */
fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

/* 使用指针访问值 */
fmt.Printf("*ip 变量的值: %d\n", *ip )
```

##### 局部变量&& 全局变量
> 全局变量的作用域是整个包，局部变量的作用域是该变量所在的花括号内
```
package main
import "fmt"
var x = 200 //定义全局变量
func localFunc() {
	fmt.Println(x) //调用全局变量
}
func main() {
	x := 1 //定义main 局部变量
	localFunc()    //调用全局变量
	fmt.Println(x) //调用main 的局部变量
	if true {
		x := 100       //定义if 局部变量
		fmt.Println(x) // 调用if 局部变量
	}
}
```

##### 引入包 && 别名
```
//引入fmt包
import  "fmt"
//引入fmt表，取一个别名fm，调用时可以直接使用fm.Println()
import fm "fmt" 
```

##### 递归
> 递归，就是在运行的过程中调用自己。

> 斐波那契数列
```
package main
import "fmt"
func fibonacci(n int) int {
  if n < 2 {
   return n
  }
  return fibonacci(n-2) + fibonacci(n-1)
}
func main() {
    var i int
    for i = 0; i < 10; i++ {
       fmt.Printf("%d\t", fibonacci(i))
    }
}
```


