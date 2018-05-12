#!-*- coding:utf-8 -*-
#!/usr/bin/python

import datetime
import sys


def fib1(n):
    if n == 1:
        return 1
    if n == 2:
        return 2

    return fib1(n-1)+fib1(n-2)

def fib2(n):
    if n == 1:
        return 1
    if n == 2:
        return 2
    if n>2:
        x,y=1,2
        for line in range(3,n+1):
            count = x + y
            x, y = y, count
    return count

def main():
    num = int(sys.argv[1])
	
    begin1 = datetime.datetime.now()
    res1 = fib1(num)
    end1 = datetime.datetime.now()
    elapsed1 = end1 - begin1

    begin2 = datetime.datetime.now()
    res2 = fib2(num)
    end2 = datetime.datetime.now()
    elapsed2 = end2 - begin2

    print "递归算法,num:{num},计算结果:{res},耗时:{elapsed1}ms".format(num=num,res=res1,elapsed1=elapsed1.microseconds)
    print "循环算法,num:{num},计算结果:{res},耗时:{elapsed2}ms".format(num=num,res=res2,elapsed2=elapsed2.microseconds)

if __name__ == "__main__":
    main()
