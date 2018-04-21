second lesson
=============

#### 字符串格式化

```
s := "hello, world!"
var s1 string  = "hello, world!"
// 类型推断
var s1  = "hello, world!"
```
#### 指针
```
    a := 10
    p := &a
    *p = 20
```
#### 引用传递
```
    var p string
    flag.StringVar(&p, "f", "./pointer.go","input")
```
#### Flag
```
    flag.BoolVar(&newlineMark, "n", false, "input")
    入参要用等号
    go run flag.go -n=false
```
#### 垃圾回收
+ 标记-清除
+ 标记-整理
+ 并行的标记-清除、整理

###### 分代gc
###### 内存碎片