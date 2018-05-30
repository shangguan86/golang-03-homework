package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Path []Point

func Distance(p Point, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func Sort(p Path) Path {
	fmt.Println("before sorted Path:", p)
	if len(p) < 2 {
		return p
	}

	for i := len(p) - 1; i >= 1; i-- {
		for j := 0; j < i; j++ {
			if p[j].X > p[j+1].X {
				p[j], p[j+1] = p[j+1], p[j]
			}
		}
	}
	fmt.Println("after sorted Path:", p)
	return p
}

func pathDistance(p Path) float64 {
	// TODO
	var res float64
	p = Sort(p)
	for i := 0; i < len(p)-1; i++ {
		res += Distance(p[i], p[i+1])
	}
	return res
}

func main() {
	p := Path{Point{0, 0}, {-1, 0}, {3, 4}}
	fmt.Println(pathDistance(p))
}
