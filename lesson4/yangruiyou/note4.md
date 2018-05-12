# note4 
> 2018/5/12


## 计算字符长度

str:=""计算字符长度kkgo"



len(str):计算str有多少个字节数,3*6+4=22byte
len([]rune(s)):计算字符个数(utf8),6+4=10
utf8.RuneCountInString(str):计算字符个数,6+4=10


