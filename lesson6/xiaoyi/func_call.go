/*
值传递，仅仅就是值;
引用传递，传递的是内存地址，修改后会改变内存地址对应储存的值;
 */
package main

import (
	"math"
	"fmt"
)

type Point struct {
	Y float64
	X float64
}

func Distance(p,q Point) float64 {
	return math.Hypot(q.X - p.Y,q.Y-p.Y)
}

// 面向对象编程，golang通过方法实现
func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X,q.Y-p.Y)
}

func main()  {
	p := Point{0,0}
	q := Point{3,4}
	fmt.Println(Distance(p,q))
	fmt.Println(p.Distance(q)) //面向对象方法
}