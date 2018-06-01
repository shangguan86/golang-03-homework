package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Path []Point

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func pathDistance(p Path) float64 {
	var num float64

	for i := 0; i < len(p)-1; i++ {

		num += p[i].Distance(p[i+1])

	}

	return float64(num)
}

func (p Path) PDistance() float64 {
	var num float64

	for i := 0; i < len(p)-1; i++ {

		num += p[i].Distance(p[i+1])

	}

	return float64(num)
}
func main() {
//	p := Path{Point{-1, 0}, {0, 0}, {3, 4}}
	p := Path{Point{-1, 0}, {0, 0}, {3, 4}}
	fmt.Println(pathDistance(p))
	fmt.Println(p.PDistance())
}
