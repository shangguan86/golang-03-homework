//求n个点到点之间的距离
package main

import (
	"math"
	"fmt"
)

type Point struct {
	X,Y float64
}

type Path []Point


// 函数实现
func pathDistance(p Path) float64 {
	sum := 0.0
	for n :=0;n < len(p)-1;n ++ {
		i := math.Hypot(p[n].X-p[n+1].Y,p[n].Y-p[n+1].X)
		sum += i

	}
	return sum
}

// 面向对象实现
func (p Path) Distance(q Path) float64 {
	sum := 0.0
	for n :=0;n < len(q)-1;n ++ {
		i := math.Hypot(q[n].X-p[n+1].X,q[n].Y-p[n+1].Y)
		sum += i

	}
	return sum
}

func main() {
	p := Path{Point{-1,0},{0,0},{3,4}}
	fmt.Println(pathDistance(p))
	fmt.Println(p.Distance(p))

}